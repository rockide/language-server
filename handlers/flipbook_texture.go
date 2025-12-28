package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var FlipbookTexture = &JsonHandler{
	Pattern: shared.FlipbookTextureGlob,
	Entries: []JsonEntry{
		{
			Store: stores.TerrainTexture.References,
			Path:  []shared.JsonPath{shared.JsonValue("*/atlas_tile")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTexture.References.Get()
			},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("*/flipbook_texture")},
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
