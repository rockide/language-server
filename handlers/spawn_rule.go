package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
)

func SpawnRule(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/spawn_rules/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.EntityId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:spawn_rules/description/identifier"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:permute_type/*/entity_type"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:delay_filter/identifier"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.BiomeTag,
				Path:    []core.JsonPath{core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:biome_filter/**/value")},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "has_biome_tag"
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityEvent,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/event"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/initial_event"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/event"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:herd/*/initial_event"),
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
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
					if node != nil {
						if id, ok := node.Value.(string); ok {
							return workspace.EntityEvent.Definition.Get(id)
						}
					}
					return workspace.EntityEvent.Definition.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					root := ctx.GetRootNode()
					node := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:spawn_rules", "description", "identifier"})
					if node != nil {
						if id, ok := node.Value.(string); ok {
							return workspace.EntityEvent.References.Get(id)
						}
					}
					return workspace.EntityEvent.References.Get()
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_filter/*"),
					core.JsonValue("minecraft:spawn_rules/conditions/*/minecraft:spawns_on_block_prevented_filter/*"),
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
		},
	}
}
