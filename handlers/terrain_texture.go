package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var TerrainTexture = &JsonHandler{
	Pattern: shared.TerrainTextureGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.TerrainTexture.Source,
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("texture_data/*/textures"),
				shared.JsonValue("texture_data/*/textures/path"),
				shared.JsonValue("texture_data/*/textures/variations/*/path"),
				shared.JsonValue("texture_data/*/textures/*"),
				shared.JsonValue("texture_data/*/textures/*/path"),
				shared.JsonValue("texture_data/*/textures/*/variations/*/path"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
