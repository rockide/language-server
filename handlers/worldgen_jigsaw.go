package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenJigsaw = &JsonHandler{
	Pattern: shared.WorldgenJigsawGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenJigsawId.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/description/identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsawId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsawId.Source.Get()
			},
		},
		{
			Store: stores.BiomeTag.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/biome_filters/**/value")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeTag.References.Get()
			},
		},
		{
			Store: stores.WorldgenTemplatePoolId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:jigsaw/start_pool"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/target"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/targets/*/data"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/target"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.References.Get()
			},
		},
	},
}
