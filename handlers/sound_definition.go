package handlers

import (
	"github.com/rockide/language-server/core"
)

func SoundDefinition(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/sounds/sound_definitions.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.SoundDefinitionId,
				Path:       []core.JsonPath{core.JsonKey("sound_definitions/*")},
				FilterDiff: true,
			},
			{
				Path: []core.JsonPath{
					core.JsonValue("sound_definitions/*/sounds/*"),
					core.JsonValue("sound_definitions/*/sounds/*/name"),
				},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.SoundPath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
		},
	}
}
