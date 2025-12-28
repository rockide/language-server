package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Atmosphere = &JsonHandler{
	Pattern: shared.AtmosphereGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.Atmosphere.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:atmosphere_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.Atmosphere.Source.Get()
			},
		},
	},
}
