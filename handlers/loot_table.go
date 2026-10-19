package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var LootTable = &JsonHandler{
	Pattern:   shared.LootTableGlob,
	PathStore: stores.LootTablePath,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "item"
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "loot_table"
			},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
