package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var gamemodes = []string{
	"a",
	"adventure",
	"c",
	"creative",
	"d",
	"default",
	"s",
	"spectator",
	"survival",
}

var Gamemode = &mcfunction.Spec{
	Name:        "gamemode",
	Description: "Sets a player's game mode.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "gameMode",
					Literals: gamemodes,
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "gameMode",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 2,
					},
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "player",
					Optional: true,
				},
			},
		},
	},
}
