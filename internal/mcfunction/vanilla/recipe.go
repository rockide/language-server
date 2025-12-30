package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Recipe = &mcfunction.Spec{
	Name:        "recipe",
	Description: "Unlocks recipe in the recipe book for a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "give",
					Literals: []string{"give"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "recipe",
					Tags:     []string{mcfunction.TagRecipeId},
					Literals: []string{"*"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "take",
					Literals: []string{"take"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "recipe",
					Tags:     []string{mcfunction.TagRecipeId},
					Literals: []string{"*"},
				},
			},
		},
	},
}
