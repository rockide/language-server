package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Clear = &mcfunction.Spec{
	Name:        "clear",
	Description: "Clears items from player inventory.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "itemName",
					Optional: true,
					Tags:     []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "maxCount",
					Optional: true,
				},
			},
		},
	},
}
