package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var SoundDefinition = &JsonHandler{
	Pattern: shared.SoundDefinitionGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.SoundDefinitionId,
			Path:       []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
			FilterDiff: true,
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("sound_definitions/*/sounds/*"),
				shared.JsonValue("sound_definitions/*/sounds/*/name"),
			},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundPath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
