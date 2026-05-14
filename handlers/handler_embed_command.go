package handlers

import "github.com/rockide/language-server/internal/mcfunction"

func newEmbedCommandParser(options mcfunction.ParserOptions) *mcfunction.Parser {
	parser := newCommandParser(options)
	parser.SetEscapedQuotes(true)
	return parser
}

var EmbedCommand = &CommandHandler{
	Parser:       newEmbedCommandParser(mcfunction.ParserOptions{}),
	EscapeQuotes: true,
}

var EmbedDialogueCommand = &CommandHandler{
	Parser: newEmbedCommandParser(mcfunction.ParserOptions{
		InitiatorSelector: true,
	}),
	EscapeQuotes: true,
}
