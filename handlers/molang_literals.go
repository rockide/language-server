package handlers

import "github.com/rockide/language-server/core"

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
	bindings []*core.SymbolBinding
	literals []string
}

var molangTypes = map[string]func(workspace *core.Workspace) molangValue{
	"EquipmentSlot": func(workspace *core.Workspace) molangValue {
		return molangValue{literals: equipmentSlots[:]}
	},
	"GraphicsMode": func(workspace *core.Workspace) molangValue {
		return molangValue{literals: graphicsModes[:]}
	},
	"InputMode": func(workspace *core.Workspace) molangValue {
		return molangValue{literals: inputModes[:]}
	},
	"BiomeId": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.BiomeId}}
	},
	"BiomeTag": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.BiomeTag}}
	},
	"BlockTag": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.BlockTag}}
	},
	"BlockAndItemTag": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.BlockTag, workspace.ItemTag}}
	},
	"BlockState": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.BlockState}}
	},
	"EntityIdentifier": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.EntityId}}
	},
	"EntityProperty": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.EntityProperty}}
	},
	"TypeFamily": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.EntityFamily}}
	},
	"ItemIdentifier": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.ItemId}}
	},
	"ItemTag": func(workspace *core.Workspace) molangValue {
		return molangValue{bindings: []*core.SymbolBinding{workspace.ItemTag}}
	},
}
