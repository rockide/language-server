package core

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/rockide/vanilla"
)

type Workspace struct {
	RootDir   string
	Documents *textdocument.Store

	// Pack directories relative to RootDir; empty when absent.
	BehaviorPack string
	ResourcePack string

	// BP symbols
	AimAssistId            *SymbolBinding
	AimAssistCategory      *SymbolBinding
	AnimationId            *SymbolBinding
	AnimationAlias         *SymbolBinding
	BlockCustomComponent   *SymbolBinding
	BiomeId                *SymbolBinding
	BiomeTag               *SymbolBinding
	BlockState             *SymbolBinding
	BlockTag               *SymbolBinding
	CameraId               *SymbolBinding
	ControllerState        *SymbolBinding
	CooldownCategory       *SymbolBinding
	DialogueId             *SymbolBinding
	EntityId               *SymbolBinding
	EntityProperty         *SymbolBinding
	EntityPropertyValue    *SymbolBinding
	EntityComponentGroup   *SymbolBinding
	EntityEvent            *SymbolBinding
	EntityFamily           *SymbolBinding
	FeatureId              *SymbolBinding
	FeatureRuleId          *SymbolBinding
	ItemCustomComponent    *SymbolBinding
	ItemId                 *SymbolBinding
	ItemTag                *SymbolBinding
	ProvidedFogId          *SymbolBinding
	RecipeId               *SymbolBinding
	RecipeTag              *SymbolBinding
	ScoreboardObjective    *SymbolBinding
	Tag                    *SymbolBinding
	TickingArea            *SymbolBinding
	WorldgenProcessorId    *SymbolBinding
	WorldgenTemplatePoolId *SymbolBinding
	WorldgenJigsawId       *SymbolBinding

	// RP symbols
	AtmosphereId          *SymbolBinding
	BlockCullingId        *SymbolBinding
	ColorGradingId        *SymbolBinding
	ClientAnimationId     *SymbolBinding
	ClientAnimationAlias  *SymbolBinding
	ClientControllerState *SymbolBinding
	EntityMaterial        *SymbolBinding
	FogId                 *SymbolBinding
	GeometryId            *SymbolBinding
	GeometryBone          *SymbolBinding
	ItemTextureId         *SymbolBinding
	Lang                  *SymbolBinding
	LightingId            *SymbolBinding
	ParticleId            *SymbolBinding
	ParticleEvent         *SymbolBinding
	ParticleMaterial      *SymbolBinding
	RenderControllerId    *SymbolBinding
	SoundDefinitionId     *SymbolBinding
	MusicDefinitionId     *SymbolBinding
	TerrainTextureId      *SymbolBinding
	WaterId               *SymbolBinding

	// Paths
	LootTablePath  *PathStore
	McFunctionPath *PathStore
	StructurePath  *PathStore
	TradeTablePath *PathStore
	SoundPath      *PathStore
	TexturePath    *PathStore
}

func NewWorkspace() *Workspace {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &Workspace{
		RootDir:   wd,
		Documents: textdocument.NewStore(),

		AimAssistId:            NewSymbolBinding(nil),
		AimAssistCategory:      NewSymbolBinding(nil),
		AnimationId:            NewSymbolBinding(nil),
		AnimationAlias:         NewSymbolBinding(nil),
		BlockCustomComponent:   NewSymbolBinding(nil),
		BiomeId:                NewSymbolBinding(vanilla.BiomeId),
		BiomeTag:               NewSymbolBinding(vanilla.BiomeTag),
		BlockState:             NewSymbolBinding(vanilla.BlockState),
		BlockTag:               NewSymbolBinding(vanilla.BlockTag),
		CameraId:               NewSymbolBinding(vanilla.CameraId),
		ControllerState:        NewSymbolBinding(nil),
		CooldownCategory:       NewSymbolBinding(vanilla.CooldownCategory),
		DialogueId:             NewSymbolBinding(nil),
		EntityId:               NewSymbolBinding(vanilla.EntityId),
		EntityProperty:         NewSymbolBinding(nil),
		EntityPropertyValue:    NewSymbolBinding(nil),
		EntityComponentGroup:   NewSymbolBinding(nil),
		EntityEvent:            NewSymbolBinding(nil),
		EntityFamily:           NewSymbolBinding(vanilla.TypeFamily),
		FeatureId:              NewSymbolBinding(vanilla.FeatureId),
		FeatureRuleId:          NewSymbolBinding(nil),
		ItemCustomComponent:    NewSymbolBinding(nil),
		ItemId:                 NewSymbolBinding(vanilla.ItemId.Union(vanilla.BlockId)), // Blocks are contained within the "block" scope
		ItemTag:                NewSymbolBinding(vanilla.ItemTag),
		ProvidedFogId:          NewSymbolBinding(nil),
		RecipeId:               NewSymbolBinding(vanilla.RecipeId),
		RecipeTag:              NewSymbolBinding(vanilla.RecipeTag),
		ScoreboardObjective:    NewSymbolBinding(nil),
		Tag:                    NewSymbolBinding(nil),
		TickingArea:            NewSymbolBinding(nil),
		WorldgenProcessorId:    NewSymbolBinding(vanilla.WorldgenProcessorId),
		WorldgenTemplatePoolId: NewSymbolBinding(vanilla.WorldgenTemplatePoolId),
		WorldgenJigsawId:       NewSymbolBinding(nil),

		AtmosphereId:          NewSymbolBinding(vanilla.AtmosphereId),
		BlockCullingId:        NewSymbolBinding(nil),
		ColorGradingId:        NewSymbolBinding(vanilla.ColorGradingId),
		ClientAnimationId:     NewSymbolBinding(vanilla.ClientAnimationId),
		ClientAnimationAlias:  NewSymbolBinding(nil),
		ClientControllerState: NewSymbolBinding(nil),
		EntityMaterial:        NewSymbolBinding(vanilla.EntityMaterial),
		FogId:                 NewSymbolBinding(vanilla.FogId),
		GeometryId:            NewSymbolBinding(vanilla.GeometryId),
		GeometryBone:          NewSymbolBinding(nil),
		ItemTextureId:         NewSymbolBinding(vanilla.ItemTextureId),
		Lang:                  NewSymbolBinding(vanilla.LangKey),
		LightingId:            NewSymbolBinding(vanilla.LightingId),
		ParticleId:            NewSymbolBinding(vanilla.ParticleId),
		ParticleEvent:         NewSymbolBinding(nil),
		ParticleMaterial:      NewSymbolBinding(vanilla.ParticleMaterial),
		RenderControllerId:    NewSymbolBinding(vanilla.RenderControllerId),
		SoundDefinitionId:     NewSymbolBinding(vanilla.SoundDefinitionId),
		MusicDefinitionId:     NewSymbolBinding(vanilla.MusicDefinitionId),
		TerrainTextureId:      NewSymbolBinding(vanilla.TerrainTextureId),
		WaterId:               NewSymbolBinding(vanilla.WaterId),

		LootTablePath:  NewPathStore(vanilla.LootTable),
		McFunctionPath: NewPathStore(nil, ".mcfunction"),
		StructurePath:  NewPathStore(nil, ".mcstructure"),
		TradeTablePath: NewPathStore(vanilla.TradeTable),
		SoundPath:      NewPathStore(vanilla.SoundPath, ".fsb", ".ogg", ".wav"),
		TexturePath:    NewPathStore(vanilla.TexturePath, ".png", ".tga", ".jpg", ".jpeg", ".texture_set.json"),
	}
}

func (w *Workspace) setPacks(bp string, rp string) {
	if bp != "" {
		w.BehaviorPack = filepath.ToSlash(filepath.Clean(bp))
	}
	if rp != "" {
		w.ResourcePack = filepath.ToSlash(filepath.Clean(rp))
	}
}

// LoadPacks sets the pack directories from the user settings, or detects them in
// the current or 'packs' directory.
func (w *Workspace) LoadPacks(options *RockideOptions) error {
	// try to get project paths from user settings
	if options != nil && options.ProjectPaths != nil {
		paths := options.ProjectPaths
		if paths.BehaviorPack != "" || paths.ResourcePack != "" {
			w.setPacks(paths.BehaviorPack, paths.ResourcePack)
			return nil
		}
	}

	// if not found, search the current dir and the 'packs' dir
	dir := "."
	if stat, err := os.Stat("packs"); err == nil && stat.IsDir() {
		dir = "packs"
	}
	fsys := os.DirFS(dir)

	var bp string
	bpPaths, err := doublestar.Glob(fsys, "{behavior_pack,*BP,BP_*,*bp,bp_*}", doublestar.WithFailOnIOErrors())
	if err != nil {
		return err
	}
	if len(bpPaths) > 0 {
		bp = dir + "/" + bpPaths[0]
		log.Printf("Behavior pack: %s", bp)
	}

	var rp string
	rpPaths, err := doublestar.Glob(fsys, "{resource_pack,*RP,RP_*,*rp,rp_*}", doublestar.WithFailOnIOErrors())
	if err != nil {
		return err
	}
	if len(rpPaths) > 0 {
		rp = dir + "/" + rpPaths[0]
		log.Printf("Resource pack: %s", rp)
	}

	if bp == "" && rp == "" {
		return errors.New("not a minecraft project")
	}

	w.setPacks(bp, rp)
	return nil
}
