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
			Store:      stores.SoundDefinitionId.Source,
			Path:       []shared.JsonPath{shared.JsonKey("sound_definitions/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.Source.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("sound_definitions/*/sounds/*"),
				shared.JsonValue("sound_definitions/*/sounds/*/name"),
			},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundPath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
