package handlers

import (
	"slices"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Biome = &JsonHandler{
	Pattern: shared.BiomeGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.BiomeId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:biome/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:    Source,
			Binding: stores.BiomeTag,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:biome/components/minecraft:tags/tags/*")},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BiomeTag.Source.Get(), stores.BiomeTag.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: sliceutil.Map([]string{
				"minecraft:mountain_parameters/steep_material_adjustment/material",
				"minecraft:mountain_parameters/steep_material_adjustment/material/name",
				"minecraft:surface_builder/builder/foundation_material",
				"minecraft:surface_builder/builder/foundation_material/name",
				"minecraft:surface_builder/builder/mid_material",
				"minecraft:surface_builder/builder/mid_material/name",
				"minecraft:surface_builder/builder/sea_floor_material",
				"minecraft:surface_builder/builder/sea_floor_material/name",
				"minecraft:surface_builder/builder/sea_material",
				"minecraft:surface_builder/builder/sea_material/name",
				"minecraft:surface_builder/builder/top_material",
				"minecraft:surface_builder/builder/top_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/floor_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/floor_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/foundation_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/foundation_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/mid_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/mid_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_floor_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_floor_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_material/name",
				"minecraft:surface_material_adjustments/adjustments/*/materials/sea_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/top_material",
				"minecraft:surface_material_adjustments/adjustments/*/materials/top_material/name",
			}, func(value string) shared.JsonPath {
				return shared.JsonValue("minecraft:biome/components/" + value)
			}),
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
			Binding: stores.FeatureId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:biome/components/minecraft:forced_features/*/*/places_feature/*"),
			},
		},
		// TODO: Add support for states.
	},
}
