package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Aimassist = &mcfunction.Spec{
	Name:        "aimassist",
	Description: "Enable Aim Assist",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "x angle",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "y angle",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "max distance",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target mode",
					Optional: true,
					Literals: []string{"angle", "distance"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "preset id",
					Optional: true,
					Tags:     []string{mcfunction.TagAimAssistId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "clear",
					Literals: []string{"clear"},
				},
			},
		},
	},
}
