package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Playanimation = &mcfunction.Spec{
	Name:        "playanimation",
	Description: "Makes one or more entities play a one-off animation. Assumes all variables are setup correctly.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "animation",
					Tags: []string{mcfunction.TagClientAnimationId},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "next_state",
					Optional: true,
					Tags:     []string{mcfunction.TagClientAnimationId},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "blend_out_time",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "stop_expression",
					Optional: true,
					Literals: []string{`""`},
					Tags:     []string{mcfunction.TagMolang},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "controller",
					Optional: true,
				},
			},
		},
	},
}
