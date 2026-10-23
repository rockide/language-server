package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
)

func FeatureRule(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/feature_rules/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.FeatureRuleId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:feature_rules/description/identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.FeatureId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:feature_rules/description/places_feature")},
			},
			{
				Kind:    References,
				Binding: workspace.BiomeTag,
				Path:    []core.JsonPath{core.JsonValue("minecraft:feature_rules/conditions/minecraft:biome_filter/**/value")},
				Matcher: func(ctx *JsonContext) bool {
					parent := ctx.GetParentNode()
					test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
					return test != nil && test.Value == "has_biome_tag"
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:feature_rules/distribution/iterations"),
			core.JsonValue("minecraft:feature_rules/distribution/scatter_chance"),
			core.JsonValue("minecraft:feature_rules/distribution/*/extent/*"),
			core.JsonValue("minecraft:feature_rules/distribution/x"),
			core.JsonValue("minecraft:feature_rules/distribution/y"),
			core.JsonValue("minecraft:feature_rules/distribution/z"),
		},
	}
}
