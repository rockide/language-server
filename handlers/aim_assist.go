package handlers

import (
	"github.com/rockide/language-server/core"
)

func AimAssistPreset(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/aim_assist/presets/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.AimAssistId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:aim_assist_preset/identifier")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.AimAssistCategory,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:aim_assist_preset/default_item_settings"),
					core.JsonValue("minecraft:aim_assist_preset/hand_settings"),
					core.JsonValue("minecraft:aim_assist_preset/item_settings/*"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:aim_assist_preset/exclusion_list/*"),
					core.JsonValue("minecraft:aim_assist_preset/liquid_targeting_list/*"),
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
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonKey("minecraft:aim_assist_preset/item_settings/*")},
			},
		},
	}
}

func AimAssistCategory(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/aim_assist/categories/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.AimAssistCategory,
				Path:       []core.JsonPath{core.JsonValue("minecraft:aim_assist_categories/categories/*/name")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/blocks/*")},
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
				Binding: workspace.EntityId,
				Path:    []core.JsonPath{core.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/entities/*")},
			},
		},
	}
}
