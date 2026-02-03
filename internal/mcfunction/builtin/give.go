package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Give = &mcfunction.Spec{
	Name:        "give",
	Description: "Gives an item to a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "itemName",
					Tags: []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amount",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindItemNbt,
					Name:     "components",
					Optional: true,
				},
			},
		},
	},
}
