package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenProcessor = &JsonHandler{
	Pattern: shared.WorldgenProcessorGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.WorldgenProcessorId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:processor_list/description/identifier")},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:processor_list/processors/**/rules/*/block_entity_modifier/loot_table")},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/block_state/name"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/block_state/name"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/output_state/name"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    References,
			Binding: stores.BlockTag,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/input_predicate/tag"),
				shared.JsonValue("minecraft:processor_list/processors/**/rules/*/location_predicate/tag"),
			},
		},
		// TODO: Block states
	},
}
