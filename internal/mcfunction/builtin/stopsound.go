package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Stopsound = &mcfunction.Spec{
	Name:        "stopsound",
	Description: "Stops a sound.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "sound",
					Optional: true,
					Tags:     []string{mcfunction.TagSoundId},
				},
			},
		},
	},
}
