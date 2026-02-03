package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Teleport = &mcfunction.Spec{
	Name:        "teleport",
	Aliases:     []string{"tp"},
	Description: "Teleports entities (players, mobs, etc.).",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindRelativeNumber,
					Name:     "yRot",
					Optional: true,
					Tags:     []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindRelativeNumber,
					Name:     "xRot",
					Optional: true,
					Tags:     []string{mcfunction.TagPitch},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindRelativeNumber,
					Name:     "yRot",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindRelativeNumber,
					Name:     "xRot",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "destination",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "destination",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "checkForBlocks",
					Optional: true,
				},
			},
		},
	},
}
