package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var MusicDefintion = &JsonHandler{
	Pattern: shared.MusicDefinitionGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.MusicDefinitionId,
			Path:       []shared.JsonPath{shared.JsonKey("*")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.SoundDefinitionId,
			Path:    []shared.JsonPath{shared.JsonValue("*/event_name")},
		},
	},
}
