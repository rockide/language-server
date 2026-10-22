package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
	"go.lsp.dev/uri"
)

type Path struct {
	Pattern shared.Pattern
	Store   *stores.PathStore
}

func (s *Path) GetPattern() shared.Pattern {
	return s.Pattern
}

func (s *Path) Parse(uri uri.URI) error {
	s.Store.Insert(s.Pattern, uri)
	return nil
}

func (s *Path) GetPaths() []core.Symbol {
	return s.Store.Get()
}

func (s *Path) Delete(uri uri.URI) {
	s.Store.Delete(uri)
}
