package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Fog = &mcfunction.Spec{
	Name:        "fog",
	Description: "Add or remove fog settings file",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "victim",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"push"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "fogId",
					Tags: []string{mcfunction.TagFogId},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "userProvidedId",
					Tags: []string{mcfunction.TagProvidedFogId},
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
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"pop", "remove"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "userProvidedId",
					Tags: []string{mcfunction.TagProvidedFogId},
				},
			},
		},
	},
}
