package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Testforblocks = &mcfunction.Spec{
	Name:        "testforblocks",
	Description: "Tests whether the blocks in two regions match.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "begin",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "end",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Optional: true,
					Literals: []string{"all", "masked"},
				},
			},
		},
	},
}
