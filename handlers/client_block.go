package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientBlock = &JsonHandler{
	Pattern: shared.ClientBlockGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonKey("*")},
			Matcher: func(ctx *JsonContext) bool {
				return ctx.NodeValue != "format_version"
			},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    References,
			Binding: stores.TerrainTextureId,
			Path: []shared.JsonPath{
				shared.JsonValue("*/textures"),
				shared.JsonValue("*/textures/*"),
			},
		},
	},
}
