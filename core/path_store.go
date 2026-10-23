package core

import (
	"path/filepath"
	"slices"
	"strings"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"go.lsp.dev/uri"
)

type PathStore struct {
	store        []Symbol
	trimSuffixes []string
	mu           sync.RWMutex
	VanillaData  mapset.Set[string]
}

func NewPathStore(vanillaData mapset.Set[string], trimSuffixes ...string) *PathStore {
	return &PathStore{trimSuffixes: trimSuffixes, VanillaData: vanillaData}
}

func (s *PathStore) Insert(workspace *Workspace, pattern Pattern, uri uri.URI) {
	packDir, ok := pattern.PackDir(workspace)
	if !ok {
		return
	}
	path, err := filepath.Rel(workspace.RootDir, uri.FsPath())
	if err != nil {
		panic(err)
	}
	path = filepath.ToSlash(path)
	_, path, found := strings.Cut(path, packDir+"/")
	if !found {
		panic("invalid project path")
	}
	for _, suffix := range s.trimSuffixes {
		str, ok := strings.CutSuffix(path, suffix)
		if ok {
			path = str
			break
		}
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store = append(s.store, Symbol{Value: path, URI: uri})
}

func (s *PathStore) Get() []Symbol {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return slices.Clone(s.store)
}

func (s *PathStore) Delete(uri uri.URI) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store = slices.DeleteFunc(s.store, func(symbol Symbol) bool {
		return symbol.URI == uri
	})
}
