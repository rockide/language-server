package handlers

import (
	"github.com/rockide/language-server/core"
)

func WorldgenStructureSet(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/worldgen/structure_sets/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.WorldgenJigsawId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:structure_set/structures/*/structure")},
			},
		},
	}
}
