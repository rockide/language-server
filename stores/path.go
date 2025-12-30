package stores

import (
	"path/filepath"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/vanilla"
)

type PathStore struct {
	store        []core.Symbol
	trimSuffixes []string
	VanillaData  mapset.Set[string]
}

func NewPathStore(vanillaData mapset.Set[string], trimSuffixes ...string) *PathStore {
	return &PathStore{trimSuffixes: trimSuffixes, VanillaData: vanillaData}
}

func (s *PathStore) Insert(pattern shared.Pattern, uri protocol.DocumentURI) {
	packDir, ok := pattern.PackDir()
	if !ok {
		return
	}
	path, err := filepath.Rel(shared.Getwd(), uri.Path())
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
	s.store = append(s.store, core.Symbol{Value: path, URI: uri})
}

func (s *PathStore) Get() []core.Symbol {
	return s.store
}

func (s *PathStore) Delete(uri protocol.DocumentURI) {
	s.store = slices.DeleteFunc(s.store, func(symbol core.Symbol) bool {
		return symbol.URI == uri
	})
}

var (
	LootTablePath  = NewPathStore(vanilla.LootTable)
	McFunctionPath = NewPathStore(nil, ".mcfunction")
	StructurePath  = NewPathStore(nil, ".mcstructure")
	TradeTablePath = NewPathStore(vanilla.TradeTable)
	SoundPath      = NewPathStore(vanilla.SoundPath, ".fsb", ".ogg", ".wav")
	TexturePath    = NewPathStore(vanilla.TexturePath, ".png", ".tga", ".jpg", ".jpeg", ".texture_set.json")
)
