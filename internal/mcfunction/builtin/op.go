package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Op = &mcfunction.Spec{
	Name:        "op",
	Description: "Grants operator status to a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
			},
		},
	},
}
