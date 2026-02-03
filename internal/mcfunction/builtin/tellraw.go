package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Tellraw = &mcfunction.Spec{
	Name:        "tellraw",
	Description: "Sends a JSON message to players.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindRawMessage,
					Name: "raw json message",
				},
			},
		},
	},
}
