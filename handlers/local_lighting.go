package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var LocalLighting = &JsonHandler{
	Pattern: shared.LocalLightingGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:local_light_settings/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Definition.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
	},
}
