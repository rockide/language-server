package handlers

import (
	"slices"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Camera = &JsonHandler{
	Pattern: shared.CameraGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.CameraId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/identifier")},
			FilterDiff: true,
		},
		{
			// FIXME: Prevent circular inheritance
			Kind:    References,
			Binding: stores.CameraId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/inherit_from")},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				parent := ctx.GetParentNode()
				identifier := jsonc.FindNodeAtLocation(parent, jsonc.Path{"identifier"})
				if identifier != nil {
					return slices.DeleteFunc(stores.CameraId.Definition.Get(), func(symbol core.Symbol) bool {
						return symbol.Value == identifier.Value
					})
				}
				return stores.CameraId.Definition.Get()
			},
		},
		{
			Kind:    References,
			Binding: stores.AimAssistId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:camera_preset/aim_assist/preset")},
		},
	},
}
