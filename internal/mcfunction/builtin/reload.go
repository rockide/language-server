package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Reload = &mcfunction.Spec{
	Name:        "reload",
	Description: "Reloads all function and script files from all behavior packs, or optionally reloads the world and all resource and behavior packs.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "all",
					Optional: true,
					Literals: []string{"all"},
				},
			},
		},
	},
}
