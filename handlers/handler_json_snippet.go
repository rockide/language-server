package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/textdocument"
	"go.lsp.dev/protocol"
)

type JsonSnippetHandler struct {
	*JsonHandler
	SnippetEntries []SnippetEntry
}

type SnippetEntry struct {
	Path     []core.JsonPath
	Snippets []Snippet
}

type Snippet struct {
	Label string
	Value string
}

func (j *JsonSnippetHandler) prepareSnippet(location *jsonc.Location) *SnippetEntry {
	for _, entry := range j.SnippetEntries {
		for _, jsonPath := range entry.Path {
			if jsonPath.IsKey == location.IsAtPropertyKey && location.Path.Matches(jsonPath.Path) {
				return &entry
			}
		}
	}
	return nil
}

func (j *JsonSnippetHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if location == nil {
		return nil
	}
	snippetEntry := j.prepareSnippet(location)
	if snippetEntry == nil {
		return j.JsonHandler.Completions(document, position)
	}
	result := []protocol.CompletionItem{}
	for _, snippet := range snippetEntry.Snippets {
		result = append(result, protocol.CompletionItem{
			Label:            snippet.Label,
			Kind:             protocol.CompletionItemKindSnippet,
			InsertTextFormat: protocol.InsertTextFormatSnippet,
			InsertText:       protocol.NewOptional(snippet.Value),
		})
	}
	return result
}
