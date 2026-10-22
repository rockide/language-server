package handlers

import (
	"path/filepath"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/pathutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var TextureSet = &JsonHandler{
	Pattern:   shared.TextureSetGlob,
	PathStore: stores.TexturePath,
	Entries: []JsonEntry{
		{
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:texture_set/color"),
				shared.JsonValue("minecraft:texture_set/normal"),
				shared.JsonValue("minecraft:texture_set/heightmap"),
				shared.JsonValue("minecraft:texture_set/metalness_emissive_roughness"),
				shared.JsonValue("minecraft:texture_set/metalness_emissive_roughness_subsurface"),
			},
			DisableRename: true,
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				textures := stores.TexturePath.Get()
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
