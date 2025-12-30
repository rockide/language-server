package mcfunction

var SelectorArg = map[string]ParameterSpec{
	"c": {
		Kind: ParameterKindNumber,
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
	// TODO: This
	// "has_property": {
	// Kind: ParameterKindJSON,
	// },
	// TODO: This
	// "hasitem": {
	// data: integer
	// item: string
	// location: slot location
	// quantity: integer range
	// slot: integer range
	// Kind: ParameterKindJSON,
	// },
	// TODO: This
	// "haspermission": {
	// <permission>=<enabled|disabled>
	// Kind: ParameterKindJSON,
	// },
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
	// TODO: This
	// "scores": {
	// },
	"tag": {
		Kind: ParameterKindString,
	},
	"type": {
		Kind: ParameterKindString,
		Tags: []string{TagTypeFamilyId},
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
}
