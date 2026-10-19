package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Lighting = &JsonHandler{
	Pattern: shared.LightingGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.LightingId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:lighting_settings/description/identifier")},
			FilterDiff: true,
		},
	},
}
