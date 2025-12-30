package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Daylock = &mcfunction.Spec{
	Name:        "daylock",
	Aliases:     []string{"alwaysday"},
	Description: "Locks and unlocks the day-night cycle.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "lock",
					Optional: true,
				},
			},
		},
	},
}
