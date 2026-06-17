package handlers

import (
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
)

type JsonRawMessageHandler struct {
	*JsonSnippetHandler
	CommandHandler *CommandHandler
}

func (j *JsonRawMessageHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	if location == nil {
		return nil
	}
	node := location.PreviousNode
	if location.Path.Matches(jsonc.NewPath("**/rawtext/*/score/name")) && j.CommandHandler != nil {
		result := []protocol.CompletionItem{}
		for selector, ok := range j.CommandHandler.Parser.GetSelectors() {
			if !ok {
				continue
			}
			value := `"@` + selector + `"`
			result = append(result, protocol.CompletionItem{
				Label: value,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: value,
						Range: protocol.Range{
							Start: document.PositionAt(node.Offset),
							End:   document.PositionAt(node.Offset + node.Length),
						},
					},
				},
			})
		}
		return result
	}
	return j.JsonSnippetHandler.Completions(document, position)
}
