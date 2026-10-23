package handlers

import (
	"github.com/rockide/language-server/core"
)

func FlipbookTexture(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/textures/flipbook_textures.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.TerrainTextureId,
				Path:    []core.JsonPath{core.JsonValue("*/atlas_tile")},
			},
			{
				Path:          []core.JsonPath{core.JsonValue("*/flipbook_texture")},
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
