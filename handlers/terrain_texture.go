package handlers

import (
	"github.com/rockide/language-server/core"
)

func TerrainTexture(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/textures/terrain_texture.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.TerrainTextureId,
				Path:       []core.JsonPath{core.JsonKey("texture_data/*")},
				FilterDiff: true,
			},
			{
				Path: []core.JsonPath{
					core.JsonValue("texture_data/*/textures"),
					core.JsonValue("texture_data/*/textures/path"),
					core.JsonValue("texture_data/*/textures/variations/*/path"),
					core.JsonValue("texture_data/*/textures/*"),
					core.JsonValue("texture_data/*/textures/*/path"),
					core.JsonValue("texture_data/*/textures/*/variations/*/path"),
				},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.TexturePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
		},
	}
}
