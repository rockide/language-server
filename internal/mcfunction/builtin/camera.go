package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var easeLiterals = []string{
	"linear",
	"spring",
	"in_quad",
	"out_quad",
	"in_out_quad",
	"in_cubic",
	"out_cubic",
	"in_out_cubic",
	"in_quart",
	"out_quart",
	"in_out_quart",
	"in_quint",
	"out_quint",
	"in_out_quint",
	"in_sine",
	"out_sine",
	"in_out_sine",
	"in_expo",
	"out_expo",
	"in_out_expo",
	"in_circ",
	"out_circ",
	"in_out_circ",
	"in_bounce",
	"out_bounce",
	"in_out_bounce",
	"in_back",
	"out_back",
	"in_out_back",
	"in_elastic",
	"out_elastic",
	"in_out_elastic",
}

var ease = mcfunction.ParameterSpec{
	Kind:     mcfunction.ParameterKindLiteral,
	Name:     "ease",
	Literals: easeLiterals,
}

var colorRange = &mcfunction.NumberRange{
	Min: 0,
	Max: 255,
}

var Camera = &mcfunction.Spec{
	Name:        "camera",
	Description: "Issues a camera instruction",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "default",
					Optional: true,
					Literals: []string{"default"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "attach_to_entity",
					Literals: []string{"attach_to_entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "detach_from_entity",
					Literals: []string{"detach_from_entity"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "targetEntity",
					Literals: []string{"target_entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "targetEntity",
					Literals: []string{"target_entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "targetCenterOffset",
					Literals: []string{"targetCenterOffset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xTargetCenterOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yTargetCenterOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zTargetCenterOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "removeTarget",
					Literals: []string{"remove_target"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "ease",
					Literals: []string{"ease"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "easeTime",
				},
				ease,
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "view_offset",
					Literals: []string{"view_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xViewOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yViewOffset",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity_offset",
					Literals: []string{"entity_offset"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "xEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "yEntityOffset",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "zEntityOffset",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "pos",
					Literals: []string{"pos"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "rot",
					Literals: []string{"rot"},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "xRot",
					Tags: []string{mcfunction.TagPitch},
				},
				{
					Kind: mcfunction.ParameterKindRelativeNumber,
					Name: "yRot",
					Tags: []string{mcfunction.TagYaw},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "lookAtEntity",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "facing",
					Literals: []string{"facing"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "lookAtPosition",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "set",
					Literals: []string{"set"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "preset",
					Tags: []string{mcfunction.TagAimAssistId},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "default",
					Optional: true,
					Literals: []string{"default"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "clear",
					Literals: []string{"clear"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fade",
					Literals: []string{"fade"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "time",
					Literals: []string{"time"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "fadeInSeconds",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "holdSeconds",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "fadeOutSeconds",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "color",
					Literals: []string{"color"},
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "red",
					Range: colorRange,
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "green",
					Range: colorRange,
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "blue",
					Range: colorRange,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fade",
					Literals: []string{"fade"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "time",
					Literals: []string{"time"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "fadeInSeconds",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "holdSeconds",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "fadeOutSeconds",
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fade",
					Literals: []string{"fade"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "color",
					Literals: []string{"color"},
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "red",
					Range: colorRange,
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "green",
					Range: colorRange,
				},
				{
					Kind:  mcfunction.ParameterKindInteger,
					Name:  "blue",
					Range: colorRange,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fade",
					Literals: []string{"fade"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fov_set",
					Literals: []string{"fov_set"},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "fov_value",
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "fovEaseTime",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fovEaseType",
					Optional: true,
					Literals: ease.Literals,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fov_clear",
					Literals: []string{"fov_clear"},
				},
				{
					Kind:     mcfunction.ParameterKindNumber,
					Name:     "fovEaseTime",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "fovEaseType",
					Optional: true,
					Literals: ease.Literals,
				},
			},
		},
	},
}
