package server

import (
	"context"

	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func SignatureHelp(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.SignatureHelpProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.SignatureHelp(document, params.Position), nil
}
