package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Place = &mcfunction.Spec{
	Name:        "place",
	Description: "Places a jigsaw structure, feature, or feature rule in the world.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"place"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "structure",
					Tags: []string{mcfunction.TagJigsawId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "pos",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "ignoreStartHeight",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "keepJigsaws",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeEntities",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "liquidSettings",
					Optional: true,
					Literals: []string{"apply_waterlogging", "ignore_waterlogging"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"place"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "pool",
					Tags: []string{mcfunction.TagJigsawTemplatePoolId},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "jigsawTarget",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "maxDepth",
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "pos",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "keepJigsaws",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "includeEntities",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "liquidSettings",
					Optional: true,
					Literals: []string{"apply_waterlogging", "ignore_waterlogging"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"feature"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "feature",
					Tags: []string{mcfunction.TagFeatureId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "position",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"featurerule"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "featurerule",
					Tags: []string{mcfunction.TagFeatureRuleId},
				},
				{
					Kind:     mcfunction.ParameterKindVector3,
					Name:     "position",
					Optional: true,
				},
			},
		},
	},
}
