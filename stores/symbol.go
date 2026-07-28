package stores

import (
	"slices"
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/vanilla"
)

type SymbolStore struct {
	store       map[string][]core.Symbol
	mu          sync.RWMutex
	VanillaData mapset.Set[string]
}

func NewSymbolStore(vanillaData mapset.Set[string]) *SymbolStore {
	return &SymbolStore{
		store:       make(map[string][]core.Symbol),
		VanillaData: vanillaData,
	}
}

func (s *SymbolStore) Insert(scope string, symbol core.Symbol) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[scope] = append(s.store[scope], symbol)
}

func (s *SymbolStore) Get(scopes ...string) []core.Symbol {
	s.mu.RLock()
	defer s.mu.RUnlock()

	res := []core.Symbol{}
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

func (s *SymbolStore) GetFrom(uri protocol.DocumentURI, scopes ...string) []core.Symbol {
	res := []core.Symbol{}
	for _, symbol := range s.Get(scopes...) {
		if symbol.URI == uri {
			res = append(res, symbol)
		}
	}
	return res
}

func (s *SymbolStore) Delete(uri protocol.DocumentURI) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for scope, symbols := range s.store {
		s.store[scope] = slices.DeleteFunc(symbols, func(symbol core.Symbol) bool {
			return symbol.URI == uri
		})
	}
}

type SymbolBinding struct {
	Source     *SymbolStore
	References *SymbolStore
}

func NewSymbolBinding(vanillaData mapset.Set[string]) *SymbolBinding {
	return &SymbolBinding{
		Source:     NewSymbolStore(vanillaData),
		References: NewSymbolStore(vanillaData),
	}
}

// BP
var (
	AimAssistId            = NewSymbolBinding(nil)
	AimAssistCategory      = NewSymbolBinding(nil)
	AnimationId            = NewSymbolBinding(nil)
	AnimationAlias         = NewSymbolBinding(nil)
	BlockCustomComponent   = NewSymbolBinding(nil)
	BiomeId                = NewSymbolBinding(vanilla.BiomeId)
	BiomeTag               = NewSymbolBinding(vanilla.BiomeTag)
	BlockState             = NewSymbolBinding(vanilla.BlockState)
	BlockTag               = NewSymbolBinding(vanilla.BlockTag)
	CameraId               = NewSymbolBinding(vanilla.CameraId)
	ControllerState        = NewSymbolBinding(nil)
	CooldownCategory       = NewSymbolBinding(vanilla.CooldownCategory)
	DialogueId             = NewSymbolBinding(nil)
	EntityId               = NewSymbolBinding(vanilla.EntityId)
	EntityProperty         = NewSymbolBinding(nil)
	EntityPropertyValue    = NewSymbolBinding(nil)
	EntityComponentGroup   = NewSymbolBinding(nil)
	EntityEvent            = NewSymbolBinding(nil)
	EntityFamily           = NewSymbolBinding(vanilla.TypeFamily)
	FeatureId              = NewSymbolBinding(vanilla.FeatureId)
	FeatureRuleId          = NewSymbolBinding(nil)
	ItemCustomComponent    = NewSymbolBinding(nil)
	ItemId                 = NewSymbolBinding(vanilla.ItemId.Union(vanilla.BlockId)) // Blocks are contained within the "block" scope
	ItemTag                = NewSymbolBinding(vanilla.ItemTag)
	ProvidedFogId          = NewSymbolBinding(nil)
	RecipeId               = NewSymbolBinding(vanilla.RecipeId)
	RecipeTag              = NewSymbolBinding(vanilla.RecipeTag)
	ScoreboardObjective    = NewSymbolBinding(nil)
	Tag                    = NewSymbolBinding(nil)
	TickingArea            = NewSymbolBinding(nil)
	WorldgenProcessorId    = NewSymbolBinding(vanilla.WorldgenProcessorId)
	WorldgenTemplatePoolId = NewSymbolBinding(vanilla.WorldgenTemplatePoolId)
	WorldgenJigsawId       = NewSymbolBinding(nil)
)

// RP
var (
	AtmosphereId          = NewSymbolBinding(vanilla.AtmosphereId)
	BlockCullingId        = NewSymbolBinding(nil)
	ColorGradingId        = NewSymbolBinding(vanilla.ColorGradingId)
	ClientAnimationId     = NewSymbolBinding(vanilla.ClientAnimationId)
	ClientAnimationAlias  = NewSymbolBinding(nil)
	ClientControllerState = NewSymbolBinding(nil)
	EntityMaterial        = NewSymbolBinding(vanilla.EntityMaterial)
	FogId                 = NewSymbolBinding(vanilla.FogId)
	GeometryId            = NewSymbolBinding(vanilla.GeometryId)
	GeometryBone          = NewSymbolBinding(nil)
	ItemTextureId         = NewSymbolBinding(vanilla.ItemTextureId)
	Lang                  = NewSymbolBinding(nil)
	LightingId            = NewSymbolBinding(vanilla.LightingId)
	ParticleId            = NewSymbolBinding(vanilla.ParticleId)
	ParticleEvent         = NewSymbolBinding(nil)
	ParticleMaterial      = NewSymbolBinding(vanilla.ParticleMaterial)
	RenderControllerId    = NewSymbolBinding(vanilla.RenderControllerId)
	SoundDefinitionId     = NewSymbolBinding(vanilla.SoundDefinitionId)
	MusicDefinitionId     = NewSymbolBinding(vanilla.MusicDefinitionId)
	TerrainTextureId      = NewSymbolBinding(vanilla.TerrainTextureId)
	WaterId               = NewSymbolBinding(vanilla.WaterId)
)
