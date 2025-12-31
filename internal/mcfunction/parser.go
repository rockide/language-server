package mcfunction

import (
	"fmt"

	"github.com/rockide/language-server/internal/mcfunction/lexer"
)

type Parser struct {
	lexer    *lexer.Lexer
	commands map[string]*Spec
	options  ParserOptions
}

type ParserOptions struct {
	InitiatorSelector bool // enables @initiator selector
	EducationMode     bool // enables education edition specific commands and arguments
}

func NewParser(options ParserOptions, commands ...*Spec) *Parser {
	return &Parser{
		lexer:    lexer.New([]rune{}),
		commands: make(map[string]*Spec),
		options:  options,
	}
}

func (p *Parser) RegisterCommands(specs ...*Spec) {
	for _, spec := range specs {
		p.commands[spec.Name] = spec
		for _, alias := range spec.Aliases {
			p.commands[alias] = spec
		}
	}
}

func (p *Parser) RegisteredCommands() map[string]*Spec {
	return p.commands
}

func (p *Parser) GetSelectors() map[string]bool {
	return getSelectors(p.options)
}

func (p *Parser) Parse(input []rune) (INode, error) {
	if len(p.commands) == 0 {
		return nil, fmt.Errorf("no commands registered")
	}
	p.lexer.Reset(input)
	root := &Node{}
	tokens := []lexer.Token{}
	eol := uint32(0)
	parse := func() {
		if len(tokens) == 0 {
			return
		}
		node := parseCommand(tokens, input, p.commands, p.options, eol)
		node.end = eol
		root.addChild(node)
	}
	for token := range p.lexer.Next() {
		eol = token.End
		switch token.Kind {
		case lexer.TokenWhitespace:
			continue
		case lexer.TokenComment:
			root.addChild(&Node{
				kind:  NodeKindComment,
				start: token.Start,
				end:   token.End,
			})
		case lexer.TokenNewline:
			parse()
			tokens = []lexer.Token{}
		default:
			tokens = append(tokens, token)
		}
	}
	parse()
	root.end = uint32(len(input))
	return root, nil
}

func parseCommand(tokens []lexer.Token, input []rune, commands map[string]*Spec, options ParserOptions, eol uint32) *NodeCommand {
	first := tokens[0]
	commandInput := first.Text(input)
	spec, ok := commands[commandInput]
	node := &NodeCommand{
		Node: &Node{
			kind:  NodeKindInvalidCommand,
			start: first.Start,
			end:   eol,
		},
	}
	addDefaultArgs := func() {
		for i := 1; i < len(tokens); i++ {
			node.addChild(&NodeArg{
				Node: &Node{
					kind:  NodeKindCommandArg,
					start: tokens[i].Start,
					end:   tokens[i].End,
				},
			})
		}
	}
	node.addChild(&Node{
		kind:  NodeKindCommandLit,
		start: first.Start,
		end:   first.End,
	})
	if !ok {
		node.name = commandInput
		addDefaultArgs()
		return node
	}
	if len(spec.Overloads) == 0 && len(tokens) > 1 {
		// Immediately invalidate if no overloads exist but arguments are provided
		node.name = commandInput
		node.spec = spec
		addDefaultArgs()
		return node
	}
	if len(spec.Overloads) == 0 && len(tokens) == 1 {
		node.kind = NodeKindCommand
		node.name = commandInput
		node.spec = spec
		return node
	}
	overloadStates := make([]*overloadState, len(spec.Overloads))
	for i := range spec.Overloads {
		state := &overloadState{
			spec:     spec,
			ov:       &spec.Overloads[i],
			options:  options,
			commands: commands,
			matched:  true,
			eol:      eol,
		}
		overloadStates[i] = state
		args, _ := state.parse(input, tokens[1:])
		if state.matched {
			node.kind = NodeKindCommand
			node.name = commandInput
			node.spec = spec
			node.children = make([]INode, 0, len(args))
			for _, arg := range args {
				node.addChild(arg)
			}
		}
	}
	if node.kind == NodeKindInvalidCommand {
		node.name = commandInput
		node.spec = spec
		addDefaultArgs()
	} else {
		node.overloadStates = overloadStates
	}
	return node
}
