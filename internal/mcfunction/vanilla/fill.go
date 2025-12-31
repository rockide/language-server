package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var oldBlockHandling = []string{
	"destroy",
	"hollow",
	"keep",
	"outline",
	"replace",
}

var Fill = &mcfunction.Spec{
	Name:        "fill",
	Description: "Fills all or parts of a region with a specific block.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
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
					Literals: oldBlockHandling,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
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
					Literals: oldBlockHandling,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
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
					Literals: oldBlockHandling,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "replaceTileName",
					Optional: true,
					Tags:     []string{mcfunction.TagBlockId},
				},
				{
					Kind:     mcfunction.ParameterKindMap,
					Name:     "replaceBlockStates",
					Optional: true,
					Tags:     []string{mcfunction.TagBlockState},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Tags: []string{mcfunction.TagBlockId},
					Name: "tileName",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "oldBlockHandling",
					Literals: oldBlockHandling,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "replaceTileName",
					Optional: true,
					Tags:     []string{mcfunction.TagBlockId},
				},
				{
					Kind:     mcfunction.ParameterKindMap,
					Name:     "replaceBlockStates",
					Optional: true,
					Tags:     []string{mcfunction.TagBlockState},
				},
			},
		},
	},
}
