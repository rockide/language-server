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
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonValue("**/entries/*/name")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
				return entryType != nil && entryType.Value == "item"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
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
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
