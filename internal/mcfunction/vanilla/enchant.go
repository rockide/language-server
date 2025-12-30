package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var enchantments = []string{
	"protection",
	"fire_protection",
	"feather_falling",
	"blast_protection",
	"projectile_protection",
	"thorns",
	"respiration",
	"depth_strider",
	"aqua_affinity",
	"sharpness",
	"smite",
	"bane_of_arthropods",
	"knockback",
	"fire_aspect",
	"looting",
	"efficiency",
	"silk_touch",
	"unbreaking",
	"fortune",
	"power",
	"punch",
	"flame",
	"infinity",
	"luck_of_the_sea",
	"lure",
	"frost_walker",
	"mending",
	"binding",
	"vanishing",
	"impaling",
	"riptide",
	"loyalty",
	"channeling",
	"multishot",
	"piercing",
	"quick_charge",
	"soul_speed",
	"swift_sneak",
	"wind_burst",
	"density",
	"breach",
}

var Enchant = &mcfunction.Spec{
	Name:        "enchant",
	Description: "Adds an enchantment to a player's selected item.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "enchantmentName",
					Literals: enchantments,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "level",
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "enchantmentId",
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "level",
					Optional: true,
				},
			},
		},
	},
}
