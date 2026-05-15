package handlers

import (
	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/mcfunction/builtin"
	"github.com/rockide/language-server/shared"
)

func newCommandParser(options mcfunction.ParserOptions) *mcfunction.Parser {
	return mcfunction.NewParser(options, builtin.Commands...)
}

var defaultCommandParser = newCommandParser(mcfunction.ParserOptions{})

var McFunction = &CommandHandler{
	Pattern: shared.FunctionGlob,
	Parser:  defaultCommandParser,
}
