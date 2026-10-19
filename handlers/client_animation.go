package handlers

import (
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientAnimation = &JsonHandler{
	Pattern: shared.ClientAnimationGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.ClientAnimationId,
			Path:       []shared.JsonPath{shared.JsonKey("animations/*")},
			FilterDiff: true,
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.ClientAnimationId.References.Get() {
					if strings.HasPrefix(ref.Value, "animation.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("animations/*/anim_time_update"),
		shared.JsonValue("animations/*/blend_weight"),
		shared.JsonValue("animations/*/bones/*/rotation/*"),
		shared.JsonValue("animations/*/bones/*/rotation/*/*"),
		shared.JsonValue("animations/*/bones/*/scale"),
		shared.JsonValue("animations/*/bones/*/scale/*"),
		shared.JsonValue("animations/*/bones/*/scale/*/*"),
		shared.JsonValue("animations/*/bones/*/position/*"),
		shared.JsonValue("animations/*/bones/*/position/*/*"),
		shared.JsonValue("animations/*/loop_delay"),
		shared.JsonValue("animations/*/particle_effects/*/pre_effect_script"),
		shared.JsonValue("animations/*/start_delay"),
		shared.JsonValue("animations/*/timeline/*"),
	},
}
