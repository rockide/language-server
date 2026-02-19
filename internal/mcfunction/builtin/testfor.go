package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Testfor = &mcfunction.Spec{
	Name:        "testfor",
	Description: "Counts entities (players, mobs, items, etc.) matching specified conditions.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
			},
		},
	},
}
