package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var CraftingItemCatalog = &JsonHandler{
	Pattern: shared.CraftingItemCatalogGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/group_identifier/icon"),
				shared.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/items/*"),
			},
		},
	},
}
