package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/mcfunction/builtin"
)

func newCommandParser(options mcfunction.ParserOptions) *mcfunction.Parser {
	return mcfunction.NewParser(options, builtin.Commands...)
}

var defaultCommandParser = newCommandParser(mcfunction.ParserOptions{})

func McFunction(workspace *core.Workspace) *CommandHandler {
	return newCommandHandler(workspace, core.BehaviorPattern("/functions/**/*.mcfunction"), defaultCommandParser)
}
