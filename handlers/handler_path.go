package handlers

import (
	"github.com/rockide/language-server/core"
	"go.lsp.dev/uri"
)

type PathHandler struct {
	Workspace *core.Workspace
	Pattern   core.Pattern
	Store     *core.PathStore
}

func (s *PathHandler) GetPattern() core.Pattern {
	return s.Pattern
}

func (s *PathHandler) Parse(uri uri.URI) error {
	s.Store.Insert(s.Workspace, s.Pattern, uri)
	return nil
}

func (s *PathHandler) GetPaths() []core.Symbol {
	return s.Store.Get()
}

func (s *PathHandler) Delete(uri uri.URI) {
	s.Store.Delete(uri)
}
