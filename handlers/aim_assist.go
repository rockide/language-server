package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var AimAssistPreset = &JsonHandler{
	Pattern: shared.AimAssistPresetGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.AimAssistId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:aim_assist_preset/identifier")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.AimAssistCategory,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:aim_assist_preset/default_item_settings"),
				shared.JsonValue("minecraft:aim_assist_preset/hand_settings"),
				shared.JsonValue("minecraft:aim_assist_preset/item_settings/*"),
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:aim_assist_preset/exclusion_list/*"),
				shared.JsonValue("minecraft:aim_assist_preset/liquid_targeting_list/*"),
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
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_preset/item_settings/*")},
		},
	},
}

var AimAssistCategory = &JsonHandler{
	Pattern: shared.AimAssistCategoryGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.AimAssistCategory,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:aim_assist_categories/categories/*/name")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/blocks/*")},
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
			Binding: stores.EntityId,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:aim_assist_categories/categories/*/priorities/entities/*")},
		},
	},
}
