package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Clearspawnpoint = &mcfunction.Spec{
	Name:        "clearspawnpoint",
	Description: "Removes the spawn point for a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
	},
}
