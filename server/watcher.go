package server

import (
	"log"

	"github.com/rockide/language-server/handlers"
	"go.lsp.dev/uri"
)

func onCreate(uri uri.URI) {
	handler := handlers.Find(uri)
	if handler == nil {
		return
	}
	log.Printf("create: %s", uri)
	if err := handler.Parse(uri); err != nil {
		log.Println(err)
	}
}

func onChange(uri uri.URI) {
	handler := handlers.Find(uri)
	if handler == nil {
		return
	}
	log.Printf("change: %s", uri)
	handler.Delete(uri)
	if err := handler.Parse(uri); err != nil {
		log.Println(err)
	}
}

func onDelete(uri uri.URI) {
	handler := handlers.Find(uri)
	if handler == nil {
		return
	}
	log.Printf("delete: %s", uri)
	handler.Delete(uri)
}
