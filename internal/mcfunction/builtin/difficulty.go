package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Difficulty = &mcfunction.Spec{
	Name:        "difficulty",
	Description: "Sets the difficulty level.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "difficulty",
					Literals: []string{
						"e",
						"easy",
						"n",
						"normal",
						"h",
						"hard",
						"p",
						"peaceful",
					},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "difficulty",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 3,
					},
				},
			},
		},
	},
}
