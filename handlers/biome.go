package handlers

import (
	"slices"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
)

func Biome(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/biomes/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.BiomeId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:biome/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:    Definition,
				Binding: workspace.BiomeTag,
				Path:    []core.JsonPath{core.JsonValue("minecraft:biome/components/minecraft:tags/tags/*")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.BiomeTag.Definition.Get(), workspace.BiomeTag.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
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
				}, func(value string) core.JsonPath {
					return core.JsonValue("minecraft:biome/components/" + value)
				}),
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
				Binding: workspace.FeatureId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:biome/components/minecraft:forced_features/*/*/places_feature/*"),
				},
			},
			// TODO: Add support for states.
		},
	}
}
