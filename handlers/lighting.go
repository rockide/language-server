package handlers

import (
	"github.com/rockide/language-server/core"
)

func Lighting(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/lighting/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.LightingId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:lighting_settings/description/identifier")},
				FilterDiff: true,
			},
		},
	}
}
