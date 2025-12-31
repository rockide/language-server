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

type INodeArgPair interface {
	INodeArg
	PairSpec() (ParameterSpec, bool)
}

type NodeArgPair struct {
	*NodeArg
	spec ParameterSpec
}

func (n *NodeArgPair) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *NodeArgPair) setParent(parent INode) {
	n.parent = parent
}

func (n *NodeArgPair) setIndex(index int) {
	n.index = index
}

func (n *NodeArgPair) PairSpec() (ParameterSpec, bool) {
	return n.spec, n.spec.Kind != ParameterKindUnknown
}

type INodeArgPairChild interface {
	INodeArg
	PairKind() PairKind
	PairSpec() (ParameterSpec, bool)
}

type NodeArgPairChild struct {
	*NodeArg
	pairKind PairKind
}

func (n *NodeArgPairChild) addChild(child INode) {
	child.setParent(n)
	child.setIndex(len(n.children))
	n.children = append(n.children, child)
}

func (n *NodeArgPairChild) setParent(parent INode) {
	n.parent = parent
}

func (n *NodeArgPairChild) setIndex(index int) {
	n.index = index
}

func (n *NodeArgPairChild) PairKind() PairKind {
	return n.pairKind
}

func (n *NodeArgPairChild) PairSpec() (ParameterSpec, bool) {
	if p, ok := n.parent.(*NodeArgPair); ok {
		return p.PairSpec()
	}
	return ParameterSpec{}, false
}

func createSelectorArg(input []rune, token lexer.Token) *NodeArg {
	value := []rune(token.Text(input))
	if len(value) < 2 || value[0] != '[' || value[len(value)-1] != ']' {
		return nil
	}
	result := &NodeArg{
		Node: &Node{
			kind:  NodeKindCommandArg,
			start: token.Start,
			end:   token.End,
		},
		paramKind: ParameterKindSelectorArg,
	}
	if len(value) == 2 {
		return result
	}
	startOffset := token.Start + 1
	value = value[1 : len(value)-1]
	lex := lexer.New([]rune(value))
	keyTokens := []lexer.Token{}
	var assignToken lexer.Token
	valueTokens := []lexer.Token{}
	createPair := func() {
		tKey := mergeTokens(keyTokens...)
		tValue := mergeTokens(valueTokens...)
		node := &NodeArgPair{
			NodeArg: &NodeArg{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: tKey.Start + startOffset,
					end:   tValue.End + startOffset,
				},
				paramKind: ParameterKindMapPair,
			},
		}
		key := &NodeArgPairChild{
			NodeArg: &NodeArg{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: tKey.Start + startOffset,
					end:   tKey.End + startOffset,
				},
				paramKind: ParameterKindMapPair,
			},
			pairKind: PairKindKey,
		}
		node.addChild(key)
		keyValue := key.Text(input)
		if spec, ok := SelectorArg[keyValue]; ok {
			node.spec = spec
		}
		node.addChild(&NodeArgPairChild{
			NodeArg: &NodeArg{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: assignToken.Start + startOffset,
					end:   assignToken.End + startOffset,
				},
				paramKind: ParameterKindMapPair,
			},
			pairKind: PairKindEqual,
		})
		node.addChild(&NodeArgPairChild{
			NodeArg: &NodeArg{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: tValue.Start + startOffset,
					end:   tValue.End + startOffset,
				},
				paramKind: ParameterKindMapPair,
			},
			pairKind: PairKindValue,
		})
		keyTokens = []lexer.Token{}
		valueTokens = []lexer.Token{}
		assignToken = lexer.Token{}
		result.addChild(node)
	}
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
			if t.Kind == lexer.TokenComma {
				state = 0
			} else {
				valueTokens = append(valueTokens, t)
			}
		}
	}
	if len(keyTokens) > 0 {
		createPair()
	}
	return result
}

func mergeTokens(tokens ...lexer.Token) lexer.Token {
	if len(tokens) == 0 {
		return lexer.Token{}
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
	}
}
