package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Event = &mcfunction.Spec{
	Name:        "event",
	Description: "Triggers an event for the specified object(s)",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "eventName",
					Tags: []string{mcfunction.TagEntityEvent},
				},
			},
		},
	},
}
