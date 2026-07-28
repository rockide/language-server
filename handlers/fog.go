package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Fog = &JsonHandler{
	Pattern: shared.FogGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.FogId.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:fog_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.FogId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.FogId.Source.Get()
			},
		},
	},
}
