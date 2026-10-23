package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/mcfunction"
)

func EmbedCommand(workspace *core.Workspace) *CommandHandler {
	return newCommandHandler(workspace, "", newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes: true,
	}))
}

func EmbedDialogueCommand(workspace *core.Workspace) *CommandHandler {
	return newCommandHandler(workspace, "", newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes:      true,
		InitiatorSelector: true,
	}))
}

func EmbedEventCommand(workspace *core.Workspace) *CommandHandler {
	return newCommandHandler(workspace, "", newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes: true,
		EventAlias:   true,
	}))
}
