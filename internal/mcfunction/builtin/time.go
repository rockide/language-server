package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Time = &mcfunction.Spec{
	Name:        "time",
	Description: "Changes or queries the world's game time.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "amount",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "amount",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "time",
					Literals: []string{
						"day",
						"midnight",
						"night",
						"noon",
						"sunrise",
						"sunset",
					},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"query"},
				},
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "time",
					Literals: []string{
						"day",
						"daytime",
						"gametime",
					},
				},
			},
		},
	},
}
