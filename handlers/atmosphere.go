package handlers

import (
	"github.com/rockide/language-server/core"
)

func Atmosphere(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/atmospherics/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.AtmosphereId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:atmosphere_settings/description/identifier")},
				FilterDiff: true,
			},
		},
	}
}
