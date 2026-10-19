package handlers

import (
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var FeatureRule = &JsonHandler{
	Pattern: shared.FeatureRuleGlob,
	Entries: []JsonEntry{
		{
			Kind:    Definition,
			Binding: stores.FeatureRuleId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/identifier")},
		},
		{
			Kind:    References,
			Binding: stores.FeatureId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/description/places_feature")},
		},
		{
			Kind:    References,
			Binding: stores.BiomeTag,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:feature_rules/conditions/minecraft:biome_filter/**/value")},
			Matcher: func(ctx *JsonContext) bool {
				parent := ctx.GetParentNode()
				test := jsonc.FindNodeAtLocation(parent, jsonc.Path{"test"})
				return test != nil && test.Value == "has_biome_tag"
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:feature_rules/distribution/iterations"),
		shared.JsonValue("minecraft:feature_rules/distribution/scatter_chance"),
		shared.JsonValue("minecraft:feature_rules/distribution/*/extent/*"),
		shared.JsonValue("minecraft:feature_rules/distribution/x"),
		shared.JsonValue("minecraft:feature_rules/distribution/y"),
		shared.JsonValue("minecraft:feature_rules/distribution/z"),
	},
}
