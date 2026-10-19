package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Atmosphere = &JsonHandler{
	Pattern: shared.AtmosphereGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.AtmosphereId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:atmosphere_settings/description/identifier")},
			FilterDiff: true,
		},
	},
}
