package server

import (
	"context"

	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func References(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.ReferencesProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.References(document, params.Position, params.Context.IncludeDeclaration), nil
}
