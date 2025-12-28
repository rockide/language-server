package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var CraftingItemCatalog = &JsonHandler{
	Pattern: shared.CraftingItemCatalogGlob,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/group_identifier/icon"),
				shared.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/items/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
	},
}
