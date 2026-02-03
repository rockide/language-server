package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Spawnpoint = &mcfunction.Spec{
	Name:        "spawnpoint",
	Description: "Sets the spawn point for a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPos",
					Optional: true,
				},
			},
		},
	},
}
