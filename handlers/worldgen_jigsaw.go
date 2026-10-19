package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenJigsaw = &JsonHandler{
	Pattern: shared.WorldgenJigsawGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.WorldgenJigsawId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/description/identifier")},
		},
		{
			Kind:    References,
			Binding: stores.BiomeTag,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:jigsaw/biome_filters/**/value")},
		},
		{
			Kind:    References,
			Binding: stores.WorldgenTemplatePoolId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:jigsaw/start_pool"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/target"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/targets/*/data"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/alias"),
				shared.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/target"),
			},
		},
	},
}
