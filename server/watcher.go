package server

import (
	"log"

	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
)

func onCreate(uri protocol.DocumentURI) {
	handler := handlers.Find(uri)
	if handler == nil {
		return
	}
	log.Printf("create: %s", uri)
	if err := handler.Parse(uri); err != nil {
		log.Println(err)
	}
}

func onChange(uri protocol.DocumentURI) {
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

func onDelete(uri protocol.DocumentURI) {
	handler := handlers.Find(uri)
	if handler == nil {
		return
	}
	log.Printf("delete: %s", uri)
	handler.Delete(uri)
}
