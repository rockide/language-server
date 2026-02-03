package handlers

import (
	"github.com/rockide/language-server/stores"
)

var equipmentSlots = [...]string{
	"slot.weapon.mainhand",
	"slot.weapon.offhand",
	"slot.armor.head",
	"slot.armor.chest",
	"slot.armor.legs",
	"slot.armor.feet",
	"slot.armor.body",
	"slot.hotbar",
	"slot.inventory",
	"slot.enderchest",
	"slot.saddle",
	"slot.armor",
	"slot.chest",
	"slot.equippable",
}

var graphicsModes = [...]string{"simple", "fancy", "deferred", "raytraced"}

var inputModes = [...]string{"keyboard_and_mouse", "touch", "gamepad", "motion_controller"}

type molangValue struct {
	bindings []*stores.SymbolBinding
	literals []string
}

var molangTypes = map[string]func() molangValue{
	"EquipmentSlot": func() molangValue {
		return molangValue{literals: equipmentSlots[:]}
	},
	"GraphicsMode": func() molangValue {
		return molangValue{literals: graphicsModes[:]}
	},
	"InputMode": func() molangValue {
		return molangValue{literals: inputModes[:]}
	},
	"BiomeId": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BiomeId}}
	},
	"BiomeTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BiomeTag}}
	},
	"BlockTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BlockTag}}
	},
	"BlockAndItemTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BlockTag, stores.ItemTag}}
	},
	"BlockState": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.BlockState}}
	},
	"EntityIdentifier": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityId}}
	},
	"EntityProperty": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityProperty}}
	},
	"TypeFamily": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.EntityFamily}}
	},
	"ItemIdentifier": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.ItemId}}
	},
	"ItemTag": func() molangValue {
		return molangValue{bindings: []*stores.SymbolBinding{stores.ItemTag}}
	},
}
