package handlers

import (
	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/mcfunction/builtin"
	"github.com/rockide/language-server/shared"
)

func newCommandParser(options mcfunction.ParserOptions) *mcfunction.Parser {
	parser := mcfunction.NewParser(options)
	parser.RegisterCommands(builtin.Commands...)
	return parser
}

var defaultCommandParser = newCommandParser(mcfunction.ParserOptions{})

var McFunction = &CommandHandler{
	Pattern: shared.FunctionGlob,
	Parser:  defaultCommandParser,
}
