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
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("minecraft:local_light_settings/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
	},
}
