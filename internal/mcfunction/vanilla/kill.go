package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Kill = &mcfunction.Spec{
	Name:        "kill",
	Description: "Kills entities (players, mobs, etc.).",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "target",
					Optional: true,
				},
			},
		},
	},
}
