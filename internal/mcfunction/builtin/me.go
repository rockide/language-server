package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Me = &mcfunction.Spec{
	Name:        "me",
	Description: "Displays a message about yourself.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:   mcfunction.ParameterKindString,
					Name:   "message",
					Greedy: true,
				},
			},
		},
	},
}
