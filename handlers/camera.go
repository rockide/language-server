package handlers

import (
	"slices"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
)

func Camera(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/cameras/presets/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.CameraId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:camera_preset/identifier")},
				FilterDiff: true,
			},
			{
				// FIXME: Prevent circular inheritance
				Kind:    References,
				Binding: workspace.CameraId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:camera_preset/inherit_from")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					parent := ctx.GetParentNode()
					identifier := jsonc.FindNodeAtLocation(parent, jsonc.Path{"identifier"})
					if identifier != nil {
						return slices.DeleteFunc(workspace.CameraId.Definition.Get(), func(symbol core.Symbol) bool {
							return symbol.Value == identifier.Value
						})
					}
					return workspace.CameraId.Definition.Get()
				},
			},
			{
				Kind:    References,
				Binding: workspace.AimAssistId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:camera_preset/aim_assist/preset")},
			},
		},
	}
}
