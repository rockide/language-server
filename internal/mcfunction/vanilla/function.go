package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Function = &mcfunction.Spec{
	Name:        "function",
	Description: "Runs commands found in the corresponding function file.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
	},
}
