package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var Structure = &mcfunction.Spec{
	Name:        "structure",
	Description: "Saves or loads a structure in the world.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"save"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagStructureFile},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeEntities",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeBlocks",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"delete"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagStructureFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"load"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagStructureFile},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rotation",
					Optional: true,
					Literals: []string{
						"0_degrees",
						"90_degrees",
						"180_degrees",
						"270_degrees",
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mirror",
					Optional: true,
					Literals: []string{
						"none",
						"x",
						"xz",
						"z",
					},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeEntities",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeBlocks",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "waterlogged",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "integrity",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 100,
					},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "seed",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"load"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagStructureFile},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rotation",
					Optional: true,
					Literals: []string{
						"0_degrees",
						"90_degrees",
						"180_degrees",
						"270_degrees",
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mirror",
					Optional: true,
					Literals: []string{
						"none",
						"x",
						"xz",
						"z",
					},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "animationMode",
					Optional: true,
					Literals: []string{
						"block_by_block",
						"layer_by_layer",
					},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "animationSeconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeEntities",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeBlocks",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "waterlogged",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "integrity",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 100,
					},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "seed",
					Optional: true,
				},
			},
		},
	},
}
