package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var MusicDefintion = &JsonHandler{
	Pattern: shared.MusicDefinitionGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.MusicDefinitionId.Source,
			Path:       []shared.JsonPath{shared.JsonKey("*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinitionId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinitionId.Source.Get()
			},
		},
		{
			Store: stores.SoundDefinitionId.References,
			Path:  []shared.JsonPath{shared.JsonValue("*/event_name")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.References.Get()
			},
		},
	},
}
