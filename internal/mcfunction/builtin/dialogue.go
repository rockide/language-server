package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Dialogue = &mcfunction.Spec{
	Name:        "dialogue",
	Description: "Opens NPC dialogue for a player.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "open",
					Literals: []string{"open"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "npc",
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "player",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "sceneName",
					Optional: true,
					Tags:     []string{mcfunction.TagDialogueId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "change",
					Literals: []string{"change"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "npc",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "sceneName",
					Tags: []string{mcfunction.TagDialogueId},
				},
				{
					Kind:     mcfunction.ParameterKindSelector,
					Name:     "players",
					Optional: true,
				},
			},
		},
	},
}
