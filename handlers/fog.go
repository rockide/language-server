package handlers

import (
	"github.com/rockide/language-server/core"
)

func Fog(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/fogs/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.FogId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:fog_settings/description/identifier")},
				FilterDiff: true,
			},
		},
	}
}
