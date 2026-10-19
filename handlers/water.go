package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Water = &JsonHandler{
	Pattern: shared.WaterGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.WaterId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:water_settings/description/identifier")},
			FilterDiff: true,
		},
	},
}
