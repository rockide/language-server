package core

import (
	"github.com/rockide/language-server/internal/textdocument"
	"go.lsp.dev/protocol"
)

type CompletionProvider interface {
	Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem
}

type DefinitionProvider interface {
	Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink
}

type ReferencesProvider interface {
	References(document *textdocument.TextDocument, position protocol.Position, includeDeclaration bool) []protocol.Location
}

type RenameProvider interface {
	PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder
	Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit
}

type HoverProvider interface {
	Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover
}

type SignatureHelpProvider interface {
	SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp
}

type SemanticTokenProvider interface {
	SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens
}
