package server

import (
	"context"

	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func SemanticTokens(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.SemanticTokenProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.SemanticTokens(document), nil
}
