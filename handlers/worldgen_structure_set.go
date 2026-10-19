package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenStructureSet = &JsonHandler{
	Pattern: shared.WorldgenStructureSetGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.WorldgenJigsawId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:structure_set/structures/*/structure")},
		},
	},
}
