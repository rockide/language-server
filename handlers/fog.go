package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Fog = &JsonHandler{
	Pattern: shared.FogGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.FogId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:fog_settings/description/identifier")},
			FilterDiff: true,
		},
	},
}
