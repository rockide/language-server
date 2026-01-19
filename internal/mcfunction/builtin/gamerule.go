package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var boolGameRules = []string{
	"commandblockoutput",
	"dodaylightcycle",
	"doentitydrops",
	"dofiretick",
	"recipesunlock",
	"dolimitedcrafting",
	"domobloot",
	"domobspawning",
	"dotiledrops",
	"doweathercycle",
	"drowningdamage",
	"falldamage",
	"firedamage",
	"keepinventory",
	"mobgriefing",
	"pvp",
	"showcoordinates",
	"locatorbar",
	"showdaysplayed",
	"naturalregeneration",
	"tntexplodes",
	"sendcommandfeedback",
	"doinsomnia",
	"commandblocksenabled",
	"doimmediaterespawn",
	"showdeathmessages",
	"showtags",
	"freezedamage",
	"respawnblocksexplode",
	"showbordereffect",
	"showrecipemessages",
	"projectilescanbreakblocks",
	"tntexplosiondropdecay",
}

var intGameRules = []string{
	"maxcommandchainlength",
	"randomtickspeed",
	"functioncommandlimit",
	"spawnradius",
	"playerssleepingpercentage",
}

var Gamerule = &mcfunction.Spec{
	Name:        "gamerule",
	Description: "Sets or queries a game rule value.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rule",
					Literals: boolGameRules,
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "value",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rule",
					Literals: intGameRules,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "value",
					Optional: true,
				},
			},
		},
	},
}
