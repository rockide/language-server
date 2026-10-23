package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
)

func LootTable(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/loot_tables/**/*.json"),
		PathStore: workspace.LootTablePath,
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonValue("**/entries/*/name")},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
					return entryType != nil && entryType.Value == "item"
				},
			},
			{
				Path: []core.JsonPath{core.JsonValue("**/entries/*/name")},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					entryType := jsonc.FindNodeAtLocation(parent, jsonc.Path{"type"})
					return entryType != nil && entryType.Value == "loot_table"
				},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.LootTablePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
		},
	}
}
