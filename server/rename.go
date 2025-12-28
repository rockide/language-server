package server

import (
	"context"

	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/sourcegraph/jsonrpc2"
)

func PrepareRename(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.PrepareRenameParams) (*protocol.PrepareRenamePlaceholder, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.RenameProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.PrepareRename(document, params.Position), nil
}

func Rename(ctx context.Context, conn *jsonrpc2.Conn, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	handler, ok := handlers.Find(params.TextDocument.URI).(handlers.RenameProvider)
	if !ok {
		return nil, nil
	}
	document := textdocument.Get(params.TextDocument.URI)
	return handler.Rename(document, params.Position, params.NewName), nil
}
