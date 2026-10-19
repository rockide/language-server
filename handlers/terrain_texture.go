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
			Kind:       Source,
			Binding:    stores.TerrainTextureId,
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
