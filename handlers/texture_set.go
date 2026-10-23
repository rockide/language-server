package handlers

import (
	"path/filepath"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/pathutil"
)

func TextureSet(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/textures/**/*.texture_set.json"),
		PathStore: workspace.TexturePath,
		Entries: []JsonEntry{
			{
				Path: []core.JsonPath{
					core.JsonValue("minecraft:texture_set/color"),
					core.JsonValue("minecraft:texture_set/normal"),
					core.JsonValue("minecraft:texture_set/heightmap"),
					core.JsonValue("minecraft:texture_set/metalness_emissive_roughness"),
					core.JsonValue("minecraft:texture_set/metalness_emissive_roughness_subsurface"),
				},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					textures := workspace.TexturePath.Get()
					dir := filepath.Dir(ctx.URI.FsPath())
					var result []core.Symbol
					for _, texture := range textures {
						if pathutil.InDir(dir, texture.URI.FsPath()) && !strings.HasSuffix(texture.URI.FsPath(), ".json") {
							relativePath, err := filepath.Rel(dir, texture.URI.FsPath())
							if err != nil {
								continue
							}
							relativePath = filepath.ToSlash(relativePath)
							ext := filepath.Ext(relativePath)
							texture.Value = strings.TrimSuffix(relativePath, ext)
							result = append(result, texture)
						}
					}
					return result
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
		},
	}
}
