package mcfunction

import (
	"github.com/rockide/language-server/internal/mcfunction/lexer"
)

type PairKind uint8

const (
	PairKindUnknown PairKind = iota
	PairKindKey
	PairKindEqual
	PairKindValue
)

type nodeArgPair struct {
	*nodeArg
	keySpec   *ParameterSpec
	valueSpec *ParameterSpec
}

func (n *nodeArgPair) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArgPair) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeArgPair) KeySpec() (ParameterSpec, bool) {
	return *n.keySpec, n.keySpec != nil
}

func (n *nodeArgPair) ValueSpec() (ParameterSpec, bool) {
	return *n.valueSpec, n.valueSpec != nil
}

type nodeArgPairChild struct {
	*nodeArg
	pairKind PairKind
}

func (n *nodeArgPairChild) addChild(child Node) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *nodeArgPairChild) setParent(parent Node) {
	n.parent = parent
}

func (n *nodeArgPairChild) PairKind() PairKind {
	return n.pairKind
}

func (n *nodeArgPairChild) KeySpec() (ParameterSpec, bool) {
	parent, ok := n.parent.(*nodeArgPair)
	if ok {
		if parent.keySpec != nil {
			return *parent.keySpec, true
		}
	}
	return ParameterSpec{}, false
}

func (n *nodeArgPairChild) ValueSpec() (ParameterSpec, bool) {
	parent, ok := n.parent.(*nodeArgPair)
	if ok {
		if parent.valueSpec != nil {
			return *parent.valueSpec, true
		}
	}
	return ParameterSpec{}, false
}

func (n *nodeArgPairChild) ParamSpec() (ParameterSpec, bool) {
	switch n.pairKind {
	case PairKindKey:
		return n.KeySpec()
	case PairKindValue:
		return n.ValueSpec()
	}
	return ParameterSpec{}, false
}

func (n *nodeArgPairChild) Keys() []string {
	p, ok := n.parent.(*nodeArgPair)
	if ok {
		a, ok := p.parent.(MapNode)
		if ok {
			return a.MapSpec().Keys()
		}
	}
	return nil
}

func createPairs(input []rune, token lexer.Token, spec *MapSpec) []*nodeArgPair {
	value := []rune(token.Text(input))
	startOffset := token.Start + 1
	value = value[1 : len(value)-1]
	lex := lexer.New(value)
	keyTokens := []lexer.Token{}
	var assignToken lexer.Token
	valueTokens := []lexer.Token{}
	createPair := func() *nodeArgPair {
		start := assignToken.Start + startOffset
		end := assignToken.End + startOffset
		tKey, kOk := mergeTokens(keyTokens...)
		var keyText string
		if kOk {
			start = tKey.Start + startOffset
			if assignToken.Kind == lexer.TokenUnknown {
				end = tKey.End + startOffset
			}
		}
		tValue, vOk := mergeTokens(valueTokens...)
		if vOk {
			end = tValue.End + startOffset
		}
		pairNode := &nodeArgPair{
			nodeArg: &nodeArg{
				node: &node{
					kind:  NodeKindCommandArg,
					start: start,
					end:   end,
				},
				paramKind: ParameterKindMapPair,
			},
		}
		if kOk {
			key := &nodeArgPairChild{
				nodeArg: &nodeArg{
					node: &node{
						kind:  NodeKindCommandArg,
						start: tKey.Start + startOffset,
						end:   tKey.End + startOffset,
					},
					paramKind: ParameterKindMapPair,
				},
				pairKind: PairKindKey,
			}
			pairNode.addChild(key)
			keyText = tKey.Text(value)
			if spec != nil {
				pairNode.keySpec = spec.keySpec
				if paramSpec, ok := spec.ValueSpec(keyText); ok {
					pairNode.valueSpec = paramSpec
				}
			}
		}
		if assignToken.Kind != lexer.TokenUnknown {
			pairNode.addChild(&nodeArgPairChild{
				nodeArg: &nodeArg{
					node: &node{
						kind:  NodeKindCommandArg,
						start: assignToken.Start + startOffset,
						end:   assignToken.End + startOffset,
					},
					paramKind: ParameterKindMapPair,
				},
				pairKind: PairKindEqual,
			})
		}
		if vOk {
			first := valueTokens[0]
			first.Start += startOffset
			first.End += startOffset
			if first.Kind == lexer.TokenMap || first.Kind == lexer.TokenJSON {
				if pairNode.valueSpec != nil {
					pairs := createPairs(input, first, pairNode.valueSpec.MapSpec)
					mapNode := &nodeArgMap{
						nodeArg: &nodeArg{
							node: &node{
								kind:  NodeKindCommandArg,
								start: tValue.Start + startOffset,
								end:   tValue.End + startOffset,
							},
							paramKind: pairNode.valueSpec.Kind,
						},
						mapSpec: pairNode.valueSpec.MapSpec,
					}
					pairNode.addChild(mapNode)
					for _, p := range pairs {
						mapNode.addChild(p)
					}
				}
			} else {
				pairNode.addChild(&nodeArgPairChild{
					nodeArg: &nodeArg{
						node: &node{
							kind:  NodeKindCommandArg,
							start: tValue.Start + startOffset,
							end:   tValue.End + startOffset,
						},
						paramKind: ParameterKindMapPair,
					},
					pairKind: PairKindValue,
				})
			}
		}
		keyTokens = []lexer.Token{}
		valueTokens = []lexer.Token{}
		assignToken = lexer.Token{}
		return pairNode
	}
	pairs := []*nodeArgPair{}
	state := 0
	for t := range lex.Next() {
		if t.Kind == lexer.TokenComment || t.Kind == lexer.TokenWhitespace {
			continue
		}
		switch state {
		case 0:
			switch t.Kind {
			case lexer.TokenEquals:
				assignToken = t
				state = 1
			case lexer.TokenComma:
				continue
			default:
				keyTokens = append(keyTokens, t)
			}
		case 1:
			switch t.Kind {
			case lexer.TokenComma, lexer.TokenWhitespace:
				state = 0
				pairs = append(pairs, createPair())
			case lexer.TokenMap, lexer.TokenJSON:
				valueTokens = append(valueTokens, t)
				if len(valueTokens) == 1 {
					state = 0
					pairs = append(pairs, createPair())
				}
			default:
				valueTokens = append(valueTokens, t)
			}
		}
	}
	if assignToken.Kind != lexer.TokenUnknown || len(keyTokens) > 0 {
		pairs = append(pairs, createPair())
	}
	return pairs
}

func mergeTokens(tokens ...lexer.Token) (lexer.Token, bool) {
	if len(tokens) == 0 {
		return lexer.Token{}, false
	}
	start := tokens[0].Start
	end := tokens[0].End
	for i := 1; i < len(tokens); i++ {
		if tokens[i].Start < start {
			start = tokens[i].Start
		}
		if tokens[i].End > end {
			end = tokens[i].End
		}
	}
	return lexer.Token{
		Kind:  lexer.TokenString,
		Start: start,
		End:   end,
	}, true
}
