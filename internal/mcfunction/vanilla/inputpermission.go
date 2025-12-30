package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var permission = []string{
	"camera",
	"dismount",
	"jump",
	"lateral_movement",
	"mount",
	"move_backward",
	"move_forward",
	"move_left",
	"move_right",
	"movement",
	"sneak",
}

var Inputpermission = &mcfunction.Spec{
	Name:        "inputpermission",
	Description: "Sets whether or not a player's input can affect their character.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "option",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "targets",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "permission",
					Literals: permission,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "state",
					Literals: []string{"enabled", "disabled"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "option",
					Literals: []string{"query"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "targets",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "permission",
					Literals: permission,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "state",
					Optional: true,
					Literals: []string{"enabled", "disabled"},
				},
			},
		},
	},
}
