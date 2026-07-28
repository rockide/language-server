package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenStructureSet = &JsonHandler{
	Pattern: shared.WorldgenStructureSetGlob,
	Entries: []JsonEntry{
		{
			Store: stores.WorldgenJigsawId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:structure_set/structures/*/structure")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsawId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenJigsawId.References.Get()
			},
		},
	},
}
