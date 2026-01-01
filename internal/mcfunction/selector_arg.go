package mcfunction

var permissionValueSpec = &ParameterSpec{
	Kind: ParameterKindLiteral,
	Literals: []string{
		"enabled",
		"disabled",
	},
}

var SelectorArg = NewMapSpec(map[string]*ParameterSpec{
	"c": {
		Kind: ParameterKindInteger,
	},
	"dx": {
		Kind: ParameterKindRelativeNumber,
	},
	"dy": {
		Kind: ParameterKindRelativeNumber,
	},
	"dz": {
		Kind: ParameterKindRelativeNumber,
	},
	"family": {
		Kind: ParameterKindString,
		Tags: []string{TagTypeFamilyId},
	},
	"has_property": {
		Kind: ParameterKindMapJSON,
		MapSpec: NewMapValueSpec(nil, &ParameterSpec{
			Kind: ParameterKindBoolean,
		}),
	},
	"hasitem": {
		Kind: ParameterKindMapJSON,
		MapSpec: NewMapSpec(map[string]*ParameterSpec{
			"data": {
				Kind: ParameterKindInteger,
			},
			"item": {
				Kind: ParameterKindString,
				Tags: []string{TagItemId},
			},
			"location": {
				Kind: ParameterKindLiteral,
				Literals: []string{
					"slot.armor",
					"slot.armor.body",
					"slot.armor.chest",
					"slot.armor.feet",
					"slot.armor.head",
					"slot.armor.legs",
					"slot.chest",
					"slot.enderchest",
					"slot.equippable",
					"slot.hotbar",
					"slot.inventory",
					"slot.saddle",
					"slot.weapon.mainhand",
					"slot.weapon.offhand",
				},
			},
			"quantity": {
				Kind: ParameterKindRange,
			},
			"slot": {
				Kind: ParameterKindRange,
			},
		}),
	},
	"haspermission": {
		Kind: ParameterKindMapJSON,
		MapSpec: NewMapSpec(map[string]*ParameterSpec{
			"camera":           permissionValueSpec,
			"dismount":         permissionValueSpec,
			"jump":             permissionValueSpec,
			"lateral_movement": permissionValueSpec,
			"mount":            permissionValueSpec,
			"move_backward":    permissionValueSpec,
			"move_forward":     permissionValueSpec,
			"move_left":        permissionValueSpec,
			"move_right":       permissionValueSpec,
			"movement":         permissionValueSpec,
			"sneak":            permissionValueSpec,
		}),
	},
	"l": {
		Kind: ParameterKindInteger,
	},
	"lm": {
		Kind: ParameterKindInteger,
	},
	"m": {
		Kind: ParameterKindLiteral,
		Literals: []string{
			"a",
			"adventure",
			"c",
			"creative",
			"d",
			"default",
			"s",
			"spectator",
			"survival",
			"!a",
			"!adventure",
			"!c",
			"!creative",
			"!d",
			"!default",
			"!s",
			"!spectator",
			"!survival",
		},
	},
	"name": {
		Kind: ParameterKindString,
	},
	"r": {
		Kind: ParameterKindNumber,
	},
	"rm": {
		Kind: ParameterKindNumber,
	},
	"rx": {
		Kind: ParameterKindRelativeNumber,
	},
	"rxm": {
		Kind: ParameterKindRelativeNumber,
	},
	"ry": {
		Kind: ParameterKindRelativeNumber,
	},
	"rym": {
		Kind: ParameterKindRelativeNumber,
	},
	"scores": {
		Kind: ParameterKindMapJSON,
		MapSpec: NewMapValueSpec(nil, &ParameterSpec{
			Kind: ParameterKindRange,
		}),
	},
	"tag": {
		Kind: ParameterKindString,
	},
	"type": {
		Kind: ParameterKindString,
		Tags: []string{TagEntityId},
	},
	"x": {
		Kind: ParameterKindRelativeNumber,
	},
	"y": {
		Kind: ParameterKindRelativeNumber,
	},
	"z": {
		Kind: ParameterKindRelativeNumber,
	},
})
