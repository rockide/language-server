package vanilla

import (
	"math"

	"github.com/rockide/language-server/internal/mcfunction"
)

var Tickingarea = &mcfunction.Spec{
	Name:        "tickingarea",
	Description: "Add, remove, or list ticking areas.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
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
					Kind:     mcfunction.ParameterKindString,
					Name:     "name",
					Optional: true,
					Tags:     []string{mcfunction.TagTickingAreaId},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "preload",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"add"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "circle",
					Literals: []string{"circle"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "center",
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "radius",
					Range: &mcfunction.NumberRange{
						Min: 0,
						Max: math.MaxFloat64,
					},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "name",
					Optional: true,
					Tags:     []string{mcfunction.TagTickingAreaId},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "preload",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"remove"},
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
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"remove"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagTickingAreaId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"remove_all"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"list"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "all-dimensions",
					Optional: true,
					Literals: []string{"all-dimensions"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"preload"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "preload",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "mode",
					Literals: []string{"preload"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "name",
					Tags: []string{mcfunction.TagTickingAreaId},
				},
				{
					Kind:     mcfunction.ParameterKindBoolean,
					Name:     "preload",
					Optional: true,
				},
			},
		},
	},
}
