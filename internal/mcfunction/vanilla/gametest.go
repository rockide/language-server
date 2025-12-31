package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Gametest = &mcfunction.Spec{
	Name:        "gametest",
	Description: "Interacts with gametest.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"runthis"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"run"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "testName",
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "rotationSteps",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"run"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "testName",
				},
				{
					Kind: mcfunction.ParameterKindBoolean,
					Name: "stopOnFailure",
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "repeatCount",
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "rotationSteps",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"runset"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "tag",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "rotationSteps",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"runsetuntilfail"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "tag",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "rotationSteps",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"clearall"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"pos"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"create"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "testName",
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "width",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "height",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "depth",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"runthese"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"stopall"},
				},
			},
		},
	},
}
