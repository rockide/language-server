package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Setmaxplayers = &mcfunction.Spec{
	Name:        "setmaxplayers",
	Description: "Sets the maximum number of players for this game session.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "maxPlayers",
				},
			},
		},
	},
}
