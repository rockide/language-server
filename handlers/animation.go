package handlers

import (
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Animation = &JsonHandler{
	Pattern: shared.AnimationGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.AnimationId,
			Path:       []shared.JsonPath{shared.JsonKey("animations/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.AnimationId.References.Get() {
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
		shared.JsonValue("animations/*/timeline/*/*"),
	},
	CommandEntry: JsonCommandEntry{
		Handler:   EmbedEventCommand,
		SlashMode: SlashAlways,
		Path:      []shared.JsonPath{shared.JsonValue("animations/*/timeline/*/*")},
	},
}
