package core

import (
	"slices"
	"sync"

	"go.lsp.dev/uri"
)

type SymbolStore struct {
	store map[string][]Symbol
	mu    sync.RWMutex
}

func NewSymbolStore() *SymbolStore {
	return &SymbolStore{
		store: make(map[string][]Symbol),
	}
}

func (s *SymbolStore) Insert(scope string, symbol Symbol) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[scope] = append(s.store[scope], symbol)
}

func (s *SymbolStore) Get(scopes ...string) []Symbol {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := []Symbol{}
	if len(scopes) > 0 {
		for _, scope := range scopes {
			res = append(res, s.store[scope]...)
		}
	} else {
		for _, symbols := range s.store {
			res = append(res, symbols...)
		}
	}
	return res
}

func (s *SymbolStore) GetFrom(uri uri.URI, scopes ...string) []Symbol {
	res := []Symbol{}
	for _, symbol := range s.Get(scopes...) {
		if symbol.URI == uri {
			res = append(res, symbol)
		}
	}
	return res
}

func (s *SymbolStore) Delete(uri uri.URI) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for scope, symbols := range s.store {
		s.store[scope] = slices.DeleteFunc(symbols, func(symbol Symbol) bool {
			return symbol.URI == uri
		})
	}
}
