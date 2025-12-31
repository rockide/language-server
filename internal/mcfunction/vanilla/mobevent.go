package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Mobevent = &mcfunction.Spec{
	Name:        "mobevent",
	Description: "Controls what mob events are allowed to run.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "event",
					Literals: []string{
						"events_enabled",
						"minecraft:ender_dragon_event",
						"minecraft:pillager_patrols_event",
						"minecraft:wandering_trader_event",
					},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "value",
					Optional: true,
				},
			},
		},
	},
}
