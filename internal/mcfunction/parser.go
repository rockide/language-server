package mcfunction

import (
	"fmt"

	"github.com/rockide/language-server/internal/mcfunction/lexer"
)

type Parser struct {
	commands map[string]*Spec
	options  ParserOptions
}

type ParserOptions struct {
	InitiatorSelector bool // enables @initiator selector
	EducationMode     bool // enables education edition specific commands and arguments
	EscapeQuotes      bool // enables escaping quotes with backslash in arguments
	EventAlias        bool // enables alias event parsing
}

func NewParser(options ParserOptions, commands ...*Spec) *Parser {
	parser := &Parser{
		commands: make(map[string]*Spec),
		options:  options,
	}
	parser.RegisterCommands(commands...)
	return parser
}

func (p *Parser) EscapeQuotes() bool {
	return p.options.EscapeQuotes
}

func (p *Parser) QuickEvent() bool {
	return p.options.EventAlias
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

func (p *Parser) Parse(input []rune) (Node, error) {
	if len(p.commands) == 0 && !p.options.EventAlias {
		return nil, fmt.Errorf("no commands registered")
	}
	lex := lexer.New(input)
	lex.SetEscapedQuotes(p.options.EscapeQuotes)
	root := &node{
		kind: NodeKindFile,
	}
	tokens := []lexer.Token{}
	eol := uint32(0)
	parse := func() {
		if len(tokens) == 0 {
			return
		}
		var node *nodeCommand
		if p.options.EventAlias && tokens[0].Kind == lexer.TokenSelector {
			node = parseEventAlias(tokens, input)
		} else {
			node = parseCommand(tokens, input, p.commands, p.options, eol)
		}
		node.end = eol
		root.addChild(node)
	}
	for token := range lex.Next() {
		eol = token.End
		switch token.Kind {
		case lexer.TokenWhitespace:
			continue
		case lexer.TokenComment:
			root.addChild(&node{
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

func parseEventAlias(tokens []lexer.Token, input []rune) *nodeCommand {
	cmdNode := &nodeCommand{
		node: &node{
			kind:  NodeKindInvalidCommand,
			start: tokens[0].Start,
		},
		overloadStates: make([]*overloadState, len(eventAliasSpec.Overloads)),
	}
	state := &overloadState{
		spec:    eventAliasSpec,
		ov:      &eventAliasSpec.Overloads[0],
		matched: true,
	}
	args, _ := state.parse(input, tokens)
	if state.matched {
		cmdNode.kind = NodeKindCommand
		cmdNode.spec = eventAliasSpec
		cmdNode.children = make([]Node, 0, len(args))
		for _, arg := range args {
			cmdNode.addChild(arg)
		}
	}
	cmdNode.overloadStates[0] = state
	return cmdNode
}

func parseCommand(tokens []lexer.Token, input []rune, commands map[string]*Spec, options ParserOptions, eol uint32) *nodeCommand {
	startIndex := 1
	first := tokens[0]
	commandInput := first.Text(input)
	if commandInput[0] == '/' {
		commandInput = commandInput[1:]
	}
	spec, ok := commands[commandInput]
	cmdNode := &nodeCommand{
		node: &node{
			kind:  NodeKindInvalidCommand,
			start: first.Start,
			end:   eol,
		},
	}
	addDefaultArgs := func() {
		for i := startIndex; i < len(tokens); i++ {
			cmdNode.addChild(&nodeArg{
				node: &node{
					kind:  NodeKindCommandArg,
					start: tokens[i].Start,
					end:   tokens[i].End,
				},
			})
		}
	}
	cmdNode.addChild(&node{
		kind:  NodeKindCommandLit,
		start: first.Start,
		end:   first.End,
	})
	if !ok {
		cmdNode.name = commandInput
		addDefaultArgs()
		return cmdNode
	}
	if len(spec.Overloads) == 0 && len(tokens) > startIndex {
		// Immediately invalidate if no overloads exist but arguments are provided
		cmdNode.name = commandInput
		cmdNode.spec = spec
		addDefaultArgs()
		return cmdNode
	}
	if len(spec.Overloads) == 0 && len(tokens) == startIndex {
		cmdNode.kind = NodeKindCommand
		cmdNode.name = commandInput
		cmdNode.spec = spec
		return cmdNode
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
		args, _ := state.parse(input, tokens[startIndex:])
		if state.matched {
			cmdNode.kind = NodeKindCommand
			cmdNode.name = commandInput
			cmdNode.spec = spec
			cmdNode.children = make([]Node, 0, len(args))
			for _, arg := range args {
				cmdNode.addChild(arg)
			}
		}
	}
	if cmdNode.kind == NodeKindInvalidCommand {
		cmdNode.name = commandInput
		cmdNode.spec = spec
		addDefaultArgs()
	} else {
		cmdNode.overloadStates = overloadStates
	}
	return cmdNode
}
