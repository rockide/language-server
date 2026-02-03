package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Xp = &mcfunction.Spec{
	Name:        "xp",
	Description: "Adds or removes player experience.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "amount",
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:   mcfunction.ParameterKindSuffixedInteger,
					Name:   "amount",
					Suffix: "L",
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
	},
}
