package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ItemTexture = &JsonHandler{
	Pattern: shared.ItemTextureGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.ItemTextureId,
			Path:       []shared.JsonPath{shared.JsonKey("texture_data/*")},
			FilterDiff: true,
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("texture_data/*/textures"),
				shared.JsonValue("texture_data/*/textures/path"),
				shared.JsonValue("texture_data/*/textures/*"),
				shared.JsonValue("texture_data/*/textures/*/path"),
			},
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
