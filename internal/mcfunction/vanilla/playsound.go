package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Playsound = &mcfunction.Spec{
	Name:        "playsound",
	Description: "Plays a sound.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "sound",
					Tags: []string{mcfunction.TagSoundId},
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "position",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "volume",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "pitch",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "minimumVolume",
					Optional: true,
				},
			},
		},
	},
}
