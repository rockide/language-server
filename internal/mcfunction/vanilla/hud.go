package vanilla

import "github.com/rockide/language-server/internal/mcfunction"

var hudElements = []string{
	"air_bubbles",
	"all",
	"armor",
	"crosshair",
	"health",
	"horse_health",
	"hotbar",
	"hunger",
	"item_text",
	"paperdoll",
	"progress_bar",
	"status_effects",
	"tooltips",
	"touch_controls",
}

var Hud = &mcfunction.Spec{
	Name:        "hud",
	Description: "Changes the visibility of hud elements.",
	Overloads: []mcfunction.SpecOverload{
		{
			Parameters: []mcfunction.ParameterSpec{
				{
					Kind: mcfunction.ParameterKindSelector,
					Name: "target",
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "visible",
					Literals: []string{"hide", "reset"},
				},
				{
					Kind:     mcfunction.ParameterKindLiteral,
					Name:     "hud_element",
					Optional: true,
					Literals: hudElements,
				},
			},
		},
	},
}
