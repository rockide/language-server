package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var Titleraw = &mcfunction.Spec{
	Name:        "titleraw",
	Description: "Controls screen titles with JSON messages.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "clear",
					Literals: []string{"clear"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "reset",
					Literals: []string{"reset"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "titleLocation",
					Literals: []string{"title", "subtitle", "actionbar"},
				},
				{
					Kind: mcfunction.ParameterKindRawMessage,
					Name: "raw json titleText",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "times",
					Literals: []string{"times"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "fadeIn",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "stay",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "fadeOut",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
			},
		},
	},
}
