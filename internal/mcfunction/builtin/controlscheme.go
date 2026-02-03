package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Controlscheme = &mcfunction.Spec{
	Name:        "controlscheme",
	Description: "Sets or clears control scheme.",
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
					Kind: mcfunction.ParameterKindLiteral,
					Name: "control scheme",
					Literals: []string{
						"camera_relative",
						"camera_relative_strafe",
						"player_relative",
						"player_relative_strafe",
						"locked_player_relative_strafe",
					},
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
