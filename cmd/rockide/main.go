package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/rockide/language-server/server"
	"github.com/sourcegraph/jsonrpc2"
)

type rockideHandler struct {
	jsonrpc2.Handler
}

func (h rockideHandler) Handle(ctx context.Context, conn *jsonrpc2.Conn, req *jsonrpc2.Request) {
	switch req.Method {
	case "textDocument/didOpen",
		"textDocument/didChange",
		"textDocument/didSave",
		"textDocument/didClose":
		h.Handler.Handle(ctx, conn, req)
	default:
		go h.Handler.Handle(ctx, conn, req)
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

	ctx := context.Background()
	handler := jsonrpc2.HandlerWithError(server.Handle)
	stream := jsonrpc2.NewBufferedStream(&stdio{}, jsonrpc2.VSCodeObjectCodec{})
	<-jsonrpc2.NewConn(ctx, stream, rockideHandler{handler}).DisconnectNotify()
}
