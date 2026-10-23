package handlers

import (
	"github.com/rockide/language-server/core"
)

func WorldgenProcessor(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/worldgen/processors/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.WorldgenProcessorId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:processor_list/description/identifier")},
			},
			{
				Path:          []core.JsonPath{core.JsonValue("minecraft:processor_list/processors/**/rules/*/block_entity_modifier/loot_table")},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.LootTablePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state/name"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state/name"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state/name"),
				},
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
			{
				Kind:    References,
				Binding: workspace.BlockTag,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/tag"),
					core.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/tag"),
				},
			},
			// TODO: Block states
		},
	}
}
