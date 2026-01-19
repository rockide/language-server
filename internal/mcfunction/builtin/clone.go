package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Clone = &mcfunction.Spec{
	Name:        "clone",
	Description: "Clones blocks from one region to another.",
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
					Name:     "maskMode",
					Optional: true,
					Literals: []string{"replace", "masked"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cloneMode",
					Optional: true,
					Literals: []string{"normal", "force", "move"},
				},
			},
		},
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
					Name:     "maskMode",
					Literals: []string{"replace", "masked"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cloneMode",
					Literals: []string{"normal", "force", "move"},
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
