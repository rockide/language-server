package handlers

import (
	"github.com/rockide/language-server/core"
)

func CraftingItemCatalog(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/item_catalog/crafting_item_catalog.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/group_identifier/icon"),
					core.JsonValue("minecraft:crafting_items_catalog/categories/*/groups/*/items/*"),
				},
			},
		},
	}
}
