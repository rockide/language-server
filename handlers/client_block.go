package handlers

import (
	"github.com/rockide/language-server/core"
)

func ClientBlock(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/blocks.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonKey("*")},
				Matcher: func(ctx *JsonContext) bool {
					return ctx.NodeValue != "format_version"
				},
				FilterDiff: true,
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
			{
				Kind:    References,
				Binding: workspace.TerrainTextureId,
				Path: []core.JsonPath{
					core.JsonValue("*/textures"),
					core.JsonValue("*/textures/*"),
				},
			},
		},
	}
}
