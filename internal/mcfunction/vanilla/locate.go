package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Locate = &mcfunction.Spec{
	Name:        "locate",
	Description: "Displays the coordinates for the closest structure or biome of a given type.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"structure"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "structure",
					Tags: []string{mcfunction.TagJigsawId},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "useNewChunksOnly",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "subcommand",
					Literals: []string{"biome"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "biome",
					Tags: []string{mcfunction.TagBiomeId},
				},
			},
		},
	},
}
