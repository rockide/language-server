package handlers

import (
	"github.com/rockide/language-server/core"
)

func MusicDefintion(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/sounds/music_definitions.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.MusicDefinitionId,
				Path:       []core.JsonPath{core.JsonKey("*")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.SoundDefinitionId,
				Path:    []core.JsonPath{core.JsonValue("*/event_name")},
			},
		},
	}
}
