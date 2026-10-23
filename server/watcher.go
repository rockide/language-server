package server

import (
	"log"

	"go.lsp.dev/uri"
)

func (s *Server) onCreate(uri uri.URI) {
	handler := s.findHandler(uri)
	if handler == nil {
		return
	}
	log.Printf("create: %s", uri)
	if err := handler.Parse(uri); err != nil {
		log.Println(err)
	}
}

func (s *Server) onChange(uri uri.URI) {
	handler := s.findHandler(uri)
	if handler == nil {
		return
	}
	log.Printf("change: %s", uri)
	handler.Delete(uri)
	if err := handler.Parse(uri); err != nil {
		log.Println(err)
	}
}

func (s *Server) onDelete(uri uri.URI) {
	handler := s.findHandler(uri)
	if handler == nil {
		return
	}
	log.Printf("delete: %s", uri)
	handler.Delete(uri)
}
