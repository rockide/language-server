package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Tell = &mcfunction.Spec{
	Name:        "tell",
	Description: "Sends a private message to one or more players.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind:   mcfunction.ParameterKindString,
					Name:   "message",
					Greedy: true,
				},
			},
		},
	},
}
