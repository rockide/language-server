package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Testforblock = &mcfunction.Spec{
	Name:        "testforblock",
	Description: "Tests whether a certain block is in a specific location.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "tileName",
					Tags: []string{mcfunction.TagBlockId},
				},
				{
					Kind:     mcfunction.ParameterKindMap,
					Name:     "blockStates",
					Optional: true,
					Tags:     []string{mcfunction.TagBlockState},
				},
			},
		},
	},
}
