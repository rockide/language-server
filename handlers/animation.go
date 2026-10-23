package handlers

import (
	"strings"

	"github.com/rockide/language-server/core"
)

func Animation(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/animations/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.AnimationId,
				Path:       []core.JsonPath{core.JsonKey("animations/*")},
				FilterDiff: true,
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					filtered := []core.Symbol{}
					for _, ref := range workspace.AnimationId.References.Get() {
						if strings.HasPrefix(ref.Value, "animation.") {
							filtered = append(filtered, ref)
						}
					}
					return filtered
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("animations/*/anim_time_update"),
			core.JsonValue("animations/*/timeline/*/*"),
		},
		CommandEntry: JsonCommandEntry{
			Handler:   EmbedEventCommand(workspace),
			SlashMode: SlashAlways,
			Path:      []core.JsonPath{core.JsonValue("animations/*/timeline/*/*")},
		},
	}
}
