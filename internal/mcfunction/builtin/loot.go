package builtin

import "github.com/rockide/language-server/internal/mcfunction"

var slotTypes = []string{
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
}

var Loot = &mcfunction.Spec{
	Name:        "loot",
	Description: "Drops the given loot table into the specified inventory or into the world.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"spawn"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"spawn"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"spawn"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"give"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"give"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"give"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "players",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"insert"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"insert"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"insert"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "position",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "entity",
					Literals: []string{"entity"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
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
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Name:     "source",
					Literals: []string{"loot"},
				},
				{
					Kind: mcfunction.ParameterKindString,
					Name: "loot_table",
					Tags: []string{mcfunction.TagLootTableFile},
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Tags:     []string{mcfunction.TagItemId},
					Literals: []string{"mainhand", "offhand"},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Name:     "source",
					Literals: []string{"kill"},
				},
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "entity",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Kind: mcfunction.ParameterKindInteger,
					Name: "count",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "target",
					Literals: []string{"replace"},
				},
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
					Name:     "source",
					Literals: []string{"mine"},
				},
				{
					Kind: mcfunction.ParameterKindVector3,
					Name: "TargetBlockPosition",
				},
				{
					Kind:     mcfunction.ParameterKindString,
					Name:     "<tool>|mainhand|offhand",
					Optional: true,
					Literals: []string{"mainhand", "offhand"},
					Tags:     []string{mcfunction.TagItemId},
				},
			},
		},
	},
}
