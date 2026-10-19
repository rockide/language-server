package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Recipe = &JsonHandler{
	Pattern: shared.RecipeGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.RecipeId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_shaped/description/identifier"),
				shared.JsonValue("minecraft:recipe_shapeless/description/identifier"),
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_furnace/input"),
				shared.JsonValue("minecraft:recipe_furnace/output"),
				shared.JsonValue("minecraft:recipe_shaped/result/item"),
				shared.JsonValue("minecraft:recipe_shaped/key/*/item"),
				shared.JsonValue("minecraft:recipe_shaped/unlock/*/item"),
				shared.JsonValue("minecraft:recipe_shapeless/result/item"),
				shared.JsonValue("minecraft:recipe_shapeless/ingredients/*/item"),
				shared.JsonValue("minecraft:recipe_shapeless/unlock/*/item"),
				shared.JsonValue("minecraft:recipe_brewing_mix/input"),
				shared.JsonValue("minecraft:recipe_brewing_mix/reagent"),
				shared.JsonValue("minecraft:recipe_brewing_mix/output"),
				shared.JsonValue("minecraft:recipe_brewing_container/input"),
				shared.JsonValue("minecraft:recipe_brewing_container/reagent"),
				shared.JsonValue("minecraft:recipe_brewing_container/output"),
			},
		},
		{
			Kind:    References,
			Binding: stores.RecipeTag,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_shaped/tags/*"),
				shared.JsonValue("minecraft:recipe_shapeless/tags/*"),
				shared.JsonValue("minecraft:recipe_furnace/tags"),
				shared.JsonValue("minecraft:recipe_brewing_container/tags"),
				shared.JsonValue("minecraft:recipe_brewing_container_mix/tags"),
				shared.JsonValue("minecraft:recipe_smithing_transform/tags"),
				shared.JsonValue("minecraft:recipe_smithing_trim/tags"),
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemTag,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:recipe_shaped/key/*/tag"),
				shared.JsonValue("minecraft:recipe_shaped/unlock/*/tag"),
				shared.JsonValue("minecraft:recipe_shapeless/ingredients/*/tag"),
				shared.JsonValue("minecraft:recipe_shapeless/unlock/*/tag"),
			},
		},
	},
}
