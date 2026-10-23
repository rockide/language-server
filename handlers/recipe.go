package handlers

import (
	"github.com/rockide/language-server/core"
)

func Recipe(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/recipes/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.RecipeId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:recipe_shaped/description/identifier"),
					core.JsonValue("minecraft:recipe_shapeless/description/identifier"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:recipe_furnace/input"),
					core.JsonValue("minecraft:recipe_furnace/output"),
					core.JsonValue("minecraft:recipe_shaped/result/item"),
					core.JsonValue("minecraft:recipe_shaped/key/*/item"),
					core.JsonValue("minecraft:recipe_shaped/unlock/*/item"),
					core.JsonValue("minecraft:recipe_shapeless/result/item"),
					core.JsonValue("minecraft:recipe_shapeless/ingredients/*/item"),
					core.JsonValue("minecraft:recipe_shapeless/unlock/*/item"),
					core.JsonValue("minecraft:recipe_brewing_mix/input"),
					core.JsonValue("minecraft:recipe_brewing_mix/reagent"),
					core.JsonValue("minecraft:recipe_brewing_mix/output"),
					core.JsonValue("minecraft:recipe_brewing_container/input"),
					core.JsonValue("minecraft:recipe_brewing_container/reagent"),
					core.JsonValue("minecraft:recipe_brewing_container/output"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.RecipeTag,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:recipe_shaped/tags/*"),
					core.JsonValue("minecraft:recipe_shapeless/tags/*"),
					core.JsonValue("minecraft:recipe_furnace/tags"),
					core.JsonValue("minecraft:recipe_brewing_container/tags"),
					core.JsonValue("minecraft:recipe_brewing_container_mix/tags"),
					core.JsonValue("minecraft:recipe_smithing_transform/tags"),
					core.JsonValue("minecraft:recipe_smithing_trim/tags"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemTag,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:recipe_shaped/key/*/tag"),
					core.JsonValue("minecraft:recipe_shaped/unlock/*/tag"),
					core.JsonValue("minecraft:recipe_shapeless/ingredients/*/tag"),
					core.JsonValue("minecraft:recipe_shapeless/unlock/*/tag"),
				},
			},
		},
	}
}
