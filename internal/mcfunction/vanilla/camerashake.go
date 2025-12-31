package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var Camerashake = &mcfunction.Spec{
	Name:        "camerashake",
	Description: "Applies shaking to the players' camera with a specified intensity and duration.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "intensity",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "seconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "shakeType",
					Optional: true,
					Literals: []string{"positional", "rotational"},
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
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
	},
}
