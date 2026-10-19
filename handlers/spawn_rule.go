package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var SpawnRule = &JsonHandler{
	Pattern: shared.SpawnRuleGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.EntityId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/description/identifier"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:permute_type/*/entity_type"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:delay_filter/identifier"),
			},
		},
		{
			Kind:    References,
			Binding: stores.BiomeTag,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:biome_filter/**/value")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_biome_tag"
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityEvent,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/initial_event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/event"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/initial_event"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return id
					}
				}
				return defaultScope
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return stores.EntityEvent.Source.Get(id)
					}
				}
				return stores.EntityEvent.Source.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				root := ctx.GetRootNode()
				node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
				if node != nil {
					if id, ok := node.Value.(string); ok {
						return stores.EntityEvent.References.Get(id)
					}
				}
				return stores.EntityEvent.References.Get()
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter/*"),
				shared.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter/*"),
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
	},
}
