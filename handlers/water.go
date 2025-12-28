package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Water = &JsonHandler{
	Pattern: shared.WaterGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Water.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:water_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Water.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Water.Source.Get()
			},
		},
	},
}
