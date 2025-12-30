package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var Spreadplayers = &mcfunction.Spec{
	Name:        "spreadplayers",
	Description: "Teleports entities to random locations.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "x",
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "z",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "spreadDistance",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "maxRange",
					Range: &mcfunction.NumberRange{
						Min: 1,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "maxHeight",
					Optional: true,
				},
			},
		},
	},
}
