package handlers

import (
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/rockide/language-server/shared"
)

type JsonSnippetHandler struct {
	*JsonHandler
	SnippetEntries []SnippetEntry
}

type SnippetEntry struct {
	Path     []shared.JsonPath
	Snippets []Snippet
}

type Snippet struct {
	Label string
	Value string
}

var snippetTextFormat = protocol.SnippetTextFormat

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
			Kind:             protocol.SnippetCompletion,
			InsertTextFormat: &snippetTextFormat,
			InsertText:       snippet.Value,
		})
	}
	return result
}
