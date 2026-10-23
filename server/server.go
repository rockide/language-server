package server

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/debouncer"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type Server struct {
	protocol.UnimplementedServer

	version   string
	workspace *core.Workspace
	handlers  []handlers.Handler
	debouncer *debouncer.Debouncer[uri.URI]
}

func New(version string) *Server {
	workspace := core.NewWorkspace()

	handlers := []handlers.Handler{
		// BP
		handlers.AimAssistPreset(workspace),
		handlers.AimAssistCategory(workspace),
		handlers.Animation(workspace),
		handlers.AnimationController(workspace),
		handlers.Biome(workspace),
		handlers.Block(workspace),
		handlers.Camera(workspace),
		handlers.CraftingItemCatalog(workspace),
		handlers.Dialogue(workspace),
		handlers.Entity(workspace),
		handlers.Feature(workspace),
		handlers.FeatureRule(workspace),
		handlers.Item(workspace),
		handlers.Lang(workspace),
		handlers.LootTable(workspace),
		handlers.McFunction(workspace),
		handlers.Recipe(workspace),
		handlers.SpawnRule(workspace),
		handlers.Structure(workspace),
		handlers.TradeTable(workspace),
		handlers.WorldgenProcessor(workspace),
		handlers.WorldgenTemplatePool(workspace),
		handlers.WorldgenJigsaw(workspace),
		handlers.WorldgenStructureSet(workspace),

		// RP
		handlers.Attachable(workspace),
		handlers.Atmosphere(workspace),
		handlers.BlockCulling(workspace),
		handlers.ClientAnimation(workspace),
		handlers.ClientAnimationController(workspace),
		handlers.ClientBiome(workspace),
		handlers.ClientBlock(workspace),
		handlers.ClientEntity(workspace),
		handlers.ClientLang(workspace),
		handlers.ClientSound(workspace),
		handlers.ColorGrading(workspace),
		handlers.EntityMaterial(workspace),
		handlers.Fog(workspace),
		handlers.FlipbookTexture(workspace),
		handlers.Geometry(workspace),
		handlers.ItemTexture(workspace),
		handlers.Lighting(workspace),
		handlers.LocalLighting(workspace),
		handlers.MusicDefintion(workspace),
		handlers.Particle(workspace),
		handlers.ParticleMaterial(workspace),
		handlers.RenderController(workspace),
		handlers.Sound(workspace),
		handlers.SoundDefinition(workspace),
		handlers.TerrainTexture(workspace),
		handlers.Texture(workspace),
		handlers.TextureSet(workspace),
		handlers.Water(workspace),
	}

	return &Server{
		version:   version,
		workspace: workspace,
		handlers:  handlers,
		debouncer: debouncer.NewDebouncer[uri.URI](300 * time.Millisecond),
	}
}

func (s *Server) findHandler(uri uri.URI) handlers.Handler {
	path := filepath.ToSlash(uri.FsPath())
	for _, handler := range s.handlers {
		pattern, ok := handler.GetPattern().Resolve(s.workspace)
		if ok && doublestar.MatchUnvalidated("**/"+pattern, path) {
			return handler
		}
	}
	return nil
}

func (s *Server) indexWorkspace() {
	startTime := time.Now()
	fsys := os.DirFS(".")
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	wg.Add(len(s.handlers))
	for _, store := range s.handlers {
		go func() {
			defer wg.Done()
			pattern, ok := store.GetPattern().Resolve(s.workspace)
			if !ok {
				return
			}
			doublestar.GlobWalk(fsys, pattern, func(path string, d fs.DirEntry) error {
				if d.IsDir() {
					return nil
				}
				absPath, err := filepath.Abs(path)
				if err != nil {
					log.Printf("Error parsing file: %s\n\t%s", path, err)
					skippedFiles.Add(1)
					return nil
				}
				uri := uri.File(absPath)
				err = store.Parse(uri)
				if err != nil {
					log.Printf("Error parsing file: %s\n\t%s", path, err)
					skippedFiles.Add(1)
					return nil
				}
				totalFiles.Add(1)
				return nil
			})
		}()
	}
	wg.Wait()

	totalTime := time.Since(startTime)
	log.Printf("Scanned %d files in %s", totalFiles.Load(), totalTime)
	if count := skippedFiles.Load(); count > 0 {
		log.Printf("Skipped %d files", count)
	}
}
