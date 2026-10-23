package handlers

import (
	"github.com/rockide/language-server/core"
)

func RenderController(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/render_controllers/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.RenderControllerId,
				Path:       []core.JsonPath{core.JsonKey("render_controllers/*")},
				FilterDiff: true,
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("render_controllers/*/uv_anim/offset/*"),
			core.JsonValue("render_controllers/*/uv_anim/scale/*"),
			core.JsonValue("render_controllers/*/geometry"),
			core.JsonValue("render_controllers/*/part_visibility/*/*"),
			core.JsonValue("render_controllers/*/materials/*/*"),
			core.JsonValue("render_controllers/*/textures/*"),
			core.JsonValue("render_controllers/*/color/*"),
			core.JsonValue("render_controllers/*/overlay_color/*"),
			core.JsonValue("render_controllers/*/is_hurt_color/*"),
			core.JsonValue("render_controllers/*/on_fire_color/*"),
		},
		MolangSemanticLocations: []core.JsonPath{
			core.JsonValue("render_controllers/*/arrays/*/*/*"),
		},
	}
}
