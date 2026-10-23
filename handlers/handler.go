package handlers

import (
	"github.com/rockide/language-server/core"
	"go.lsp.dev/uri"
)

type Handler interface {
	GetPattern() core.Pattern
	Parse(uri uri.URI) error
	Delete(uri uri.URI)
}
