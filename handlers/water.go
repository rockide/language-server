package handlers

import (
	"github.com/rockide/language-server/core"
)

func Water(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/water/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.WaterId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:water_settings/description/identifier")},
				FilterDiff: true,
			},
		},
	}
}
