package handlers

import "github.com/rockide/language-server/internal/mcfunction"

var EmbedCommand = &CommandHandler{
	Parser: newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes: true,
	}),
}

var EmbedDialogueCommand = &CommandHandler{
	Parser: newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes: true,
	}),
}

var EmbedEventCommand = &CommandHandler{
	Parser: newCommandParser(mcfunction.ParserOptions{
		EscapeQuotes: true,
		EventAlias:   true,
	}),
}
