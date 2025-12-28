package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ColorGrading = &JsonHandler{
	Pattern: shared.ColorGradingGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.ColorGrading.Source,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:color_grading_settings/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGrading.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGrading.Source.Get()
			},
		},
	},
}
