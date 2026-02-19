package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Toggledownfall = &mcfunction.Spec{
	Name:        "toggledownfall",
	Description: "Toggles the weather.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{},
		},
	},
}
