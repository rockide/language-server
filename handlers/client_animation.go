package handlers

import (
	"strings"

	"github.com/rockide/language-server/core"
)

func ClientAnimation(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/animations/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.ClientAnimationId,
				Path:       []core.JsonPath{core.JsonKey("animations/*")},
				FilterDiff: true,
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					filtered := []core.Symbol{}
					for _, ref := range workspace.ClientAnimationId.References.Get() {
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
			core.JsonValue("animations/*/blend_weight"),
			core.JsonValue("animations/*/bones/*/rotation/*"),
			core.JsonValue("animations/*/bones/*/rotation/*/*"),
			core.JsonValue("animations/*/bones/*/scale"),
			core.JsonValue("animations/*/bones/*/scale/*"),
			core.JsonValue("animations/*/bones/*/scale/*/*"),
			core.JsonValue("animations/*/bones/*/position/*"),
			core.JsonValue("animations/*/bones/*/position/*/*"),
			core.JsonValue("animations/*/loop_delay"),
			core.JsonValue("animations/*/particle_effects/*/pre_effect_script"),
			core.JsonValue("animations/*/start_delay"),
			core.JsonValue("animations/*/timeline/*"),
		},
	}
}
