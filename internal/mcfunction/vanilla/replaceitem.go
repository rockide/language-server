package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var Replaceitem = &mcfunction.Spec{
	Name:        "replaceitem",
	Description: "Replaces items in inventories.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "block",
					Literals: []string{"block"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "slotType",
					Literals: []string{"slot.container"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "slotId",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "itemName",
					Tags: []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amount",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindItemNbt,
					Name:     "components",
					Optional: true,
				},
			},
		},
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
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "slotType",
					Literals: slotTypes,
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "slotId",
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "itemName",
					Tags: []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amount",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindItemNbt,
					Name:     "components",
					Optional: true,
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "block",
					Literals: []string{"block"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "slotType",
					Literals: []string{"slot.container"},
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "slotId",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "oldItemHandling",
					Literals: []string{"keep", "replace"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "itemName",
					Tags: []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amount",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindItemNbt,
					Name:     "components",
					Optional: true,
				},
			},
		},
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
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "slotType",
					Literals: slotTypes,
				},
				{
					Kind: mcfunction.ParameterKindInteger,
					Name: "slotId",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "oldItemHandling",
					Literals: []string{"keep", "replace"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "itemName",
					Tags: []string{mcfunction.TagItemId},
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "amount",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindInteger,
					Name:     "data",
					Optional: true,
				},
				{
					Kind:     mcfunction.ParameterKindItemNbt,
					Name:     "components",
					Optional: true,
				},
			},
		},
	},
}
