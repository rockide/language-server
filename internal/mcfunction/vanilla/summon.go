package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Summon = &mcfunction.Spec{
	Name:        "summon",
	Description: "Summons an entity.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPos",
					Optional: true,
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
					Kind:     mcfunction.ParameterKindString,
					Name:     "spawnEvent",
					Tags:     []string{mcfunction.TagEntityEvent},
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "nameTag",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "nameTag",
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPos",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPos",
					Optional: true,
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
					Kind:     mcfunction.ParameterKindString,
					Name:     "spawnEvent",
					Tags:     []string{mcfunction.TagEntityEvent},
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "nameTag",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "spawnPos",
					Optional: true,
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
					Kind:     mcfunction.ParameterKindString,
					Name:     "spawnEvent",
					Tags:     []string{mcfunction.TagEntityEvent},
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "nameTag",
					Optional: true,
				},
			},
		},
	},
}
