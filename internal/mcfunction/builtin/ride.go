package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Ride = &mcfunction.Spec{
	Name:        "ride",
	Description: "Makes entities ride other entities, stops entities from riding, makes rides evict their riders, or summons rides or riders.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "riders",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"start_riding"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "ride",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "teleportRules",
					Optional: true,
					Literals: []string{"teleport_ride", "teleport_rider"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "howToFill",
					Optional: true,
					Literals: []string{"if_group_fits", "until_full"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "riders",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"stop_riding"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "rides",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"evict_riders"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "rides",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"summon_rider"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "spawnEvent",
					Optional: true,
					Tags:     []string{mcfunction.TagEntityEvent},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "nameTag",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "riders",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"summon_ride"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "entityType",
					Tags: []string{mcfunction.TagEntityId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rideRules",
					Optional: true,
					Literals: []string{"no_ride_change", "reassign_rides", "skip_rides"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "spawnEvent",
					Optional: true,
					Tags:     []string{mcfunction.TagEntityEvent},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "nameTag",
					Optional: true,
				},
			},
		},
	},
}
