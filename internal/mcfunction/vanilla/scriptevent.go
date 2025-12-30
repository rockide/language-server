package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Scriptevent = &mcfunction.Spec{
	Name:        "scriptevent",
	Description: "Triggers a script event with an ID and message.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "messageId",
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
