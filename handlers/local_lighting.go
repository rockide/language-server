package handlers

import (
	"github.com/rockide/language-server/core"
)

func LocalLighting(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/local_lighting/local_lighting.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonKey("minecraft:local_light_settings/*")},
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
		},
	}
}
