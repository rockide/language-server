package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/rockide/language-server/server"
	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

func rockideHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	return func(ctx context.Context, req *jsonrpc2.Request) (any, error) {
		switch req.Method() {
		case protocol.MethodTextDocumentDidOpen,
			protocol.MethodTextDocumentDidChange,
			protocol.MethodTextDocumentDidSave,
			protocol.MethodTextDocumentDidClose:
		default:
			jsonrpc2.Async(ctx)
		}
		return handler(ctx, req)
	}
}

var version = "dev"

func main() {
	printVersion := flag.Bool("version", false, "Print version")
	flag.Parse()
	if *printVersion {
		fmt.Println(version)
		return
	}

	conn := jsonrpc2.NewConn(jsonrpc2.NewStream(&stdio{}))
	client := protocol.ClientDispatcher(conn)
	ctx := protocol.WithClient(context.Background(), client)
	handler := protocol.CancelHandler(rockideHandler(protocol.ServerHandler(server.New(version), jsonrpc2.MethodNotFoundHandler)))

	conn.Go(ctx, handler)
	<-conn.Done()
}
