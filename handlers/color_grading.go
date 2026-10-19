package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ColorGrading = &JsonHandler{
	Pattern: shared.ColorGradingGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.ColorGradingId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:color_grading_settings/description/identifier")},
			FilterDiff: true,
		},
	},
}
