package handlers

import (
	"github.com/rockide/language-server/core"
)

func ColorGrading(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/color_grading/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.ColorGradingId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:color_grading_settings/description/identifier")},
				FilterDiff: true,
			},
		},
	}
}
