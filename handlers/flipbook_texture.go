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
			Kind:    References,
			Binding: stores.TerrainTextureId,
			Path:    []shared.JsonPath{shared.JsonValue("*/atlas_tile")},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("*/flipbook_texture")},
			DisableRename: true,
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
