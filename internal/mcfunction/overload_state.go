package mcfunction

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/rockide/language-server/internal/mcfunction/lexer"
)

type overloadState struct {
	spec     *Spec
	ov       *SpecOverload
	index    int
	matched  bool // Indicates if any parameter has been matched or partially matched
	options  ParserOptions
	commands map[string]*Spec
	eol      uint32
}

func (o *overloadState) Peek() (ParameterSpec, bool) {
	if o.eof() {
		return ParameterSpec{}, false
	}
	return o.ov.Parameters[o.index], true
}

func (o *overloadState) Matched() bool {
	return o.matched
}

func (o *overloadState) Index() int {
	return o.index
}

func (o *overloadState) Parameters() []ParameterSpec {
	return o.ov.Parameters
}

func (o *overloadState) ParamAt(index int) (ParameterSpec, bool) {
	if index < 0 || index >= len(o.ov.Parameters) {
		return ParameterSpec{}, false
	}
	return o.ov.Parameters[index], true
}

func (o *overloadState) eof() bool {
	return o.index >= len(o.ov.Parameters)
}

func (o *overloadState) advance() {
	o.index++
}

func (o *overloadState) parse(input []rune, tokens []lexer.Token) ([]INodeArg, error) {
	if len(tokens) > 0 && o.eof() {
		return nil, fmt.Errorf("unexpected extra arguments")
	}
	if len(tokens) == 0 {
		if param, ok := o.Peek(); !ok || param.Optional {
			return nil, nil
		} else if ok && !param.Optional {
			return nil, fmt.Errorf("missing required arguments")
		}
	}
	args := []INodeArg{}
	tokenIndex := 0
	for tokenIndex < len(tokens) {
		param, ok := o.Peek()
		if !ok {
			o.matched = false
			return args, fmt.Errorf("excess arguments")
		}
		if param.Greedy {
			for ; tokenIndex < len(tokens); tokenIndex++ {
				arg, _, err := o.matchParameter(input, tokens, tokenIndex, param)
				if err != nil {
					break
				}
				args = append(args, arg)
				o.matched = true
			}
			break
		}

		arg, consumed, err := o.matchParameter(input, tokens, tokenIndex, param)
		if err != nil {
			o.matched = false
			break
		}
		args = append(args, arg)
		o.advance()
		tokenIndex += consumed
	}
	if !o.matched {
		return args, fmt.Errorf("no overload matched")
	}
	return args, nil
}

func (o *overloadState) matchParameter(input []rune, tokens []lexer.Token, tokenIndex int, param ParameterSpec) (INodeArg, int, error) {
	token := tokens[tokenIndex]
	text := token.Text(input)
	arg := &NodeArg{
		Node: &Node{
			kind:  NodeKindCommandArg,
			start: token.Start,
			end:   token.End,
		},
		paramKind: param.Kind,
	}
	switch param.Kind {
	case ParameterKindLiteral:
		if slices.Contains(param.Literals, token.Text(input)) {
			return arg, 1, nil
		}
	case ParameterKindString:
		return arg, 1, nil
	case ParameterKindNumber:
		if token.Kind == lexer.TokenNumber {
			return arg, 1, nil
		}
	case ParameterKindInteger:
		if token.Kind == lexer.TokenNumber {
			// Strictly integer check, not allowing decimal point
			if !strings.Contains(text, ".") {
				return arg, 1, nil
			}
		}
	case ParameterKindBoolean:
		if text == "true" || text == "false" {
			return arg, 1, nil
		}
	case ParameterKindSelector:
		if token.Kind == lexer.TokenSelector {
			selector := text[1:] // remove '@'
			validSelectors := getSelectors(o.options)
			if _, ok := validSelectors[selector]; ok {
				advance := 1
				// Check for selector arguments
				if tokenIndex+1 < len(tokens) && tokens[tokenIndex+1].Kind == lexer.TokenMap {
					next := tokens[tokenIndex+1]
					advance = 2
					selArg := createSelectorArg(input, next)
					if selArg == nil {
						arg.addChild(&NodeArg{
							Node: &Node{
								kind:  NodeKindCommandArg,
								start: next.Start,
								end:   next.End,
							},
							paramKind: ParameterKindSelectorArg,
						})
					}
					arg.addChild(selArg)
					arg.Node.end = tokens[tokenIndex+1].End
				}
				return arg, advance, nil
			}
		}
	case ParameterKindMap:
		if token.Kind == lexer.TokenMap {
			// TODO:
			return arg, 1, nil
		}
	case ParameterKindJSON:
		if token.Kind == lexer.TokenJSON {
			return arg, 1, nil
		}
	case ParameterKindVector2, ParameterKindVector3:
		size := 2
		if param.Kind == ParameterKindVector3 {
			size = 3
		}
		if tokenIndex+size-1 < len(tokens) {
			if isValidRelatives(input, tokens[tokenIndex:tokenIndex+size]...) {
				for i := range make([]int, size) {
					arg.addChild(&NodeArg{
						Node: &Node{
							kind:  NodeKindCommandArg,
							start: tokens[tokenIndex+i].Start,
							end:   tokens[tokenIndex+i].End,
						},
						paramKind: ParameterKindRelativeNumber,
					})
				}
				arg.Node.end = tokens[tokenIndex+size-1].End
				return arg, size, nil
			}
		}
	case ParameterKindRange:
		// 1. Number, Range, Number
		if matchTokens(tokens, tokenIndex, lexer.TokenNumber, lexer.TokenRange, lexer.TokenNumber) {
			arg.Node.end = tokens[tokenIndex+2].End
			return arg, 3, nil
		}
		// 2. Range, Number
		if matchTokens(tokens, tokenIndex, lexer.TokenRange, lexer.TokenNumber) {
			arg.Node.end = tokens[tokenIndex+1].End
			return arg, 2, nil
		}
		// 3. Number, Range
		if matchTokens(tokens, tokenIndex, lexer.TokenNumber, lexer.TokenRange) {
			arg.Node.end = tokens[tokenIndex+1].End
			return arg, 2, nil
		}
		// 4. Bang, Number, Range, Number
		if matchTokens(tokens, tokenIndex, lexer.TokenBang, lexer.TokenNumber, lexer.TokenRange, lexer.TokenNumber) {
			arg.Node.start = tokens[tokenIndex].Start
			arg.Node.end = tokens[tokenIndex+3].End
			return arg, 4, nil
		}
		// 5. Bang, Range, Number
		if matchTokens(tokens, tokenIndex, lexer.TokenBang, lexer.TokenRange, lexer.TokenNumber) {
			arg.Node.start = tokens[tokenIndex].Start
			arg.Node.end = tokens[tokenIndex+2].End
			return arg, 3, nil
		}
		// 6. Bang, Number, Range
		if matchTokens(tokens, tokenIndex, lexer.TokenBang, lexer.TokenNumber, lexer.TokenRange) {
			arg.Node.start = tokens[tokenIndex].Start
			arg.Node.end = tokens[tokenIndex+2].End
			return arg, 3, nil
		}
	case ParameterKindSuffixedInteger:
		if token.Kind == lexer.TokenNumber && !strings.Contains(text, ".") && tokenIndex+1 < len(tokens) {
			suffixToken := tokens[tokenIndex+1]
			suffix := param.Suffix
			suffixText := suffixToken.Text(input)
			if suffixText == suffix {
				arg.Node.end = suffixToken.End
				return arg, 2, nil
			}
		}
	case ParameterKindWildcardInteger:
		_, err := strconv.Atoi(text)
		if token.Text(input) == "*" || (err == nil && !strings.Contains(text, ".")) {
			return arg, 1, nil
		}
	case ParameterKindChainedCommand:
		arg := &NodeArgCommand{
			NodeCommand: &NodeCommand{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: token.Start,
					end:   o.eol,
				},
				name: o.spec.Name,
				spec: o.spec,
			},
			paramKind: ParameterKindChainedCommand,
		}
		rest := tokens[tokenIndex:]
		if len(rest) == 0 {
			return arg, 1, nil
		}
		overloadStates := make([]*overloadState, 0, len(o.spec.Overloads))
		for i := range o.spec.Overloads {
			state := &overloadState{
				spec:     o.spec,
				ov:       &o.spec.Overloads[i],
				options:  o.options,
				commands: o.commands,
				matched:  true,
				eol:      o.eol,
			}
			overloadStates = append(overloadStates, state)
			args, _ := state.parse(input, rest)
			if state.matched {
				arg.children = make([]INode, 0, len(args))
				for _, a := range args {
					arg.addChild(a)
				}
			}
		}
		if len(arg.children) == 0 {
			for i := range rest {
				arg.addChild(&NodeArg{
					Node: &Node{
						kind:  NodeKindCommandArg,
						start: rest[i].Start,
						end:   rest[i].End,
					},
				})
			}
		}
		arg.overloadStates = overloadStates
		return arg, math.MaxInt32, nil
	case ParameterKindCommand:
		rest := tokens[tokenIndex:]
		commandNode := parseCommand(rest, input, o.commands, o.options, o.eol)
		arg := &NodeArgCommand{
			NodeCommand: commandNode,
			paramKind:   ParameterKindCommand,
		}
		return arg, math.MaxInt32, nil
	case ParameterKindRawMessage, ParameterKindItemNbt:
		if token.Kind == lexer.TokenJSON {
			return arg, 1, nil
		}
	case ParameterKindRelativeNumber:
		if token.Kind == lexer.TokenRelativeNumber || token.Kind == lexer.TokenNumber {
			return arg, 1, nil
		}
	}
	return nil, 0, fmt.Errorf("parameter did not match")
}

func isValidRelatives(input []rune, tokens ...lexer.Token) bool {
	var b byte
	for _, token := range tokens {
		if token.Kind != lexer.TokenRelativeNumber && token.Kind != lexer.TokenNumber {
			return false
		}
		if f := token.Text(input)[0]; f == '~' || f == '^' {
			if b == 0 {
				b = f
			} else if b != f {
				return false
			}
		}
	}
	return true
}

func getSelectors(options ParserOptions) map[string]bool {
	res := map[string]bool{
		"p":         true,
		"a":         true,
		"r":         true,
		"s":         true,
		"e":         true,
		"n":         true,
		"initiator": true,
		"c":         true,
		"v":         true,
	}
	for sel := range res {
		if sel == "initiator" && !options.InitiatorSelector {
			res[sel] = false
			continue
		}
		if sel == "c" || sel == "v" {
			if !options.EducationMode {
				res[sel] = false
				continue
			}
		}
	}
	return res
}

func matchTokens(tokens []lexer.Token, index int, kinds ...lexer.TokenKind) bool {
	if index+len(kinds) > len(tokens) {
		return false
	}
	for i, kind := range kinds {
		if tokens[index+i].Kind != kind {
			return false
		}
	}
	return true
}
