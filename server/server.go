package server

import (
	"errors"
	"io/fs"
	"log"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bmatcuk/doublestar/v4"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/handlers"
	"github.com/rockide/language-server/internal/protocol"
)

func findProjectPaths(options *protocol.InitializationOptions) (*core.Project, error) {
	// try to get project paths from user settings
	if options != nil && options.ProjectPaths != nil {
		paths := options.ProjectPaths
		if paths.BehaviorPack != "" || paths.ResourcePack != "" {
			return core.NewProject(paths.BehaviorPack, paths.ResourcePack), nil
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
		return nil, err
	}
	if len(bpPaths) > 0 {
		bp = dir + "/" + bpPaths[0]
		log.Printf("Behavior pack: %s", bp)
	}

	var rp string
	rpPaths, err := doublestar.Glob(fsys, "{resource_pack,*RP,RP_*,*rp,rp_*}", doublestar.WithFailOnIOErrors())
	if err != nil {
		return nil, err
	}
	if len(rpPaths) > 0 {
		rp = dir + "/" + rpPaths[0]
		log.Printf("Resource pack: %s", rp)
	}

	if bp == "" && rp == "" {
		return nil, errors.New("not a minecraft project")
	}

	return core.NewProject(bp, rp), nil
}

func indexWorkspace() {
	startTime := time.Now()
	fsys := os.DirFS(".")
	totalFiles := atomic.Uint32{}
	skippedFiles := atomic.Uint32{}

	var wg sync.WaitGroup
	wg.Add(len(handlers.GetAll()))
	for _, store := range handlers.GetAll() {
		go func() {
			defer wg.Done()
			pattern, ok := store.GetPattern().ToString()
			if !ok {
				return
			}
			doublestar.GlobWalk(fsys, pattern, func(path string, d fs.DirEntry) error {
				if d.IsDir() {
					return nil
				}
				uri := protocol.URIFromPath(path)
				err := store.Parse(uri)
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
