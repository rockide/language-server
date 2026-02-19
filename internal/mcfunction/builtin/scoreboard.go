package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Scoreboard = &mcfunction.Spec{
	Name:        "scoreboard",
	Description: "Tracks and displays scores for various objectives.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"objectives"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "criteria",
					Literals: []string{"dummy"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "displayName",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"objectives"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"remove"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"objectives"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"list"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"objectives"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"setdisplay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "displaySlot",
					Literals: []string{"belowname", "list", "sidebar"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "objective",
					Optional: true,
					Tags:     []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "sortOrder",
					Optional: true,
					Literals: []string{"ascending", "descending"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"objectives"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"setdisplay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "displaySlot",
					Literals: []string{"belowname"},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "objective",
					Optional: true,
					Tags:     []string{mcfunction.TagScoreboardObjectiveId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"list"},
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "playername",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"reset"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "objective",
					Tags:     []string{mcfunction.TagScoreboardObjectiveId},
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"test"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind: mcfunction.ParameterKindWildcardInteger,
					Name: "min",
				},
				{
					Kind:     mcfunction.ParameterKindWildcardInteger,
					Name:     "max",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"random"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "min",
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "max",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"set", "add", "remove"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "category",
					Literals: []string{"players"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "action",
					Literals: []string{"operation"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "targetName",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "targetObjective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
				{
					Kind: mcfunction.ParameterKindLiteral,
					Name: "operation",
					Literals: []string{
						"%=",
						"*=",
						"+=",
						"-=",
						"/=",
						"<",
						"=",
						">",
						"><",
					},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "selector",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "objective",
					Tags: []string{mcfunction.TagScoreboardObjectiveId},
				},
			},
		},
	},
}
