package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var damageCauseLiterals = []string{
	"anvil",
	"block_explosion",
	"campfire",
	"charging",
	"contact",
	"drowning",
	"entity_attack",
	"entity_explosion",
	"fall",
	"falling_block",
	"fire",
	"fire_tick",
	"fireworks",
	"fly_into_wall",
	"freezing",
	"lava",
	"lightning",
	"mace_smash",
	"magic",
	"magma",
	"none",
	"override",
	"piston",
	"projectile",
	"ram_attack",
	"self_destruct",
	"sonic_boom",
	"soul_campfire",
	"stalactite",
	"stalagmite",
	"starve",
	"suffocation",
	"suicide",
	"temperature",
	"thorns",
	"void",
	"wither",
}

var Damage = &mcfunction.Spec{
	Name:        "damage",
	Description: "Apply damage to the specified entities.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "amount",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cause",
					Optional: true,
					Literals: damageCauseLiterals,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "amount",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cause",
					Literals: damageCauseLiterals,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "origin",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "damager",
				},
			},
		},
	},
}
