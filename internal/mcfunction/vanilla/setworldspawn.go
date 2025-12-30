package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Setworldspawn = &mcfunction.Spec{
	Name:        "setworldspawn",
	Description: "Sets the world spawn.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPoint",
					Optional: true,
				},
			},
		},
	},
}
