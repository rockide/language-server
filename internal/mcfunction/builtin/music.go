package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Music = &mcfunction.Spec{
	Name:        "music",
	Description: "Allows you to control playing music tracks.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"queue"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "trackName",
					Tags: []string{mcfunction.TagMusicId},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "volume",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0.01,
						Max: 1,
					},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "fadeSeconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 10,
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "repeatMode",
					Optional: true,
					Literals: []string{"loop", "play_once"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"play"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "trackName",
					Tags: []string{mcfunction.TagMusicId},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "volume",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0.01,
						Max: 1,
					},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "fadeSeconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 10,
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "repeatMode",
					Optional: true,
					Literals: []string{"loop", "play_once"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"stop"},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "fadeSeconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 10,
					},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"volume"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "volume",
					Range: &mcfunction.NumberRange{
						Min: 0.01,
						Max: 1,
					},
				},
			},
		},
	},
}
