package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Setblock = &mcfunction.Spec{
	Name:        "setblock",
	Description: "Changes a block to another block.",
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
					Kind: mcfunction.ParameterKindMap,
					Name: "blockStates",
					Tags: []string{mcfunction.TagBlockState},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "oldBlockHandling",
					Optional: true,
					Literals: []string{"destroy", "keep", "replace"},
				},
			},
		},
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
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "oldBlockHandling",
					Optional: true,
					Literals: []string{"destroy", "keep", "replace"},
				},
			},
		},
	},
}
