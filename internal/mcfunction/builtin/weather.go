package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Weather = &mcfunction.Spec{
	Name:        "weather",
	Description: "Sets the weather.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "type",
					Literals: []string{"clear", "rain", "thunder"},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "duration",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "query",
					Literals: []string{"query"},
				},
			},
		},
	},
}
