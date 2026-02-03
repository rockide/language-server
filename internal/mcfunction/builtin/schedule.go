package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var Schedule = &mcfunction.Spec{
	Name:        "schedule",
	Description: "Schedules an action to be executed once an area is loaded, or after a certain amount of time.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"delay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "time",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "DelayMode",
					Optional: true,
					Literals: []string{"append", "replace"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"delay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
				{
					Kind:   mcfunction.ParameterKindSuffixedInteger,
					Name:   "time",
					Suffix: "t",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "DelayMode",
					Literals: []string{"append", "replace"},
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"delay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
				{
					Kind:   mcfunction.ParameterKindSuffixedInteger,
					Name:   "time",
					Suffix: "s",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "DelayMode",
					Literals: []string{"append", "replace"},
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"delay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
				{
					Kind:   mcfunction.ParameterKindSuffixedInteger,
					Name:   "time",
					Suffix: "d",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "DelayMode",
					Literals: []string{"append", "replace"},
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"delay"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"clear"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"clear"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"on_area_loaded"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "from",
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "to",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"on_area_loaded"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "type",
					Literals: []string{"circle"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "center",
				},
				{
					Kind: mcfunction.ParameterKindNumber,
					Name: "radius",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"on_area_loaded"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"add"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "type",
					Literals: []string{"tickingarea"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagTickingAreaId},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"on_area_loaded"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"clear"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cleartype",
					Literals: []string{"tickingarea"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagTickingAreaId},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "function",
					Tags:     []string{mcfunction.TagFunctionFile},
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"on_area_loaded"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "condition",
					Literals: []string{"clear"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "cleartype",
					Literals: []string{"function"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "function",
					Tags: []string{mcfunction.TagFunctionFile},
				},
			},
		},
	},
}
