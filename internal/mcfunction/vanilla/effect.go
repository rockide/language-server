package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var effects = []string{
	"absorption",
	"bad_omen",
	"blindness",
	"breath_of_the_nautilus",
	"conduit_power",
	"darkness",
	"fatal_poison",
	"fire_resistance",
	"haste",
	"health_boost",
	"hunger",
	"infested",
	"instant_damage",
	"instant_health",
	"invisibility",
	"jump_boost",
	"levitation",
	"mining_fatigue",
	"nausea",
	"night_vision",
	"oozing",
	"poison",
	"raid_omen",
	"regeneration",
	"resistance",
	"saturation",
	"slow_falling",
	"slowness",
	"speed",
	"strength",
	"trial_omen",
	"village_hero",
	"water_breathing",
	"weakness",
	"weaving",
	"wind_charged",
	"wither",
}

var Effect = &mcfunction.Spec{
	Name:        "effect",
	Description: "Add or remove status effects.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "Mode",
					Literals: []string{"clear"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "effect",
					Optional: true,
					Literals: effects,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "effect",
					Literals: effects,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "seconds",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amplifier",
					Optional: true,
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: 255,
					},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "hideParticles",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "effect",
					Literals: effects,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "Mode",
					Literals: []string{"infinite"},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amplifier",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "hideParticles",
					Optional: true,
				},
			},
		},
	},
}
