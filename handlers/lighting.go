package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Lighting = &JsonHandler{
	Pattern: shared.LightingGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Lighting.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:lighting_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Lighting.Source.Get()
			},
		},
	},
}
