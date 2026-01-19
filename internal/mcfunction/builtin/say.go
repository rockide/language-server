package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Say = &mcfunction.Spec{
	Name:        "say",
	Description: "Sends a message in the chat to other players.",
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
