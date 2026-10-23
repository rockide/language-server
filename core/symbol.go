package core

import (
	mapset "github.com/deckarep/golang-set/v2"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type Symbol struct {
	Value string
	URI   uri.URI
	Range *protocol.Range
}

type SymbolBinding struct {
	Definition  *SymbolStore
	References  *SymbolStore
	VanillaData mapset.Set[string]
}

func NewSymbolBinding(vanillaData mapset.Set[string]) *SymbolBinding {
	return &SymbolBinding{
		Definition:  NewSymbolStore(),
		References:  NewSymbolStore(),
		VanillaData: vanillaData,
	}
}
