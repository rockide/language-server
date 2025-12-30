package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Tag = &mcfunction.Spec{
	Name:        "tag",
	Description: "Manages tags stored in entities.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"add", "remove"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagTagId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"list"},
				},
			},
		},
	},
}
