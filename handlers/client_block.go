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
			Store: stores.ItemId.References,
			Path:  []shared.JsonPath{shared.JsonKey("*")},
			Matcher: func(ctx *JsonContext) bool {
				return ctx.NodeValue != "format_version"
			},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Store: stores.TerrainTextureId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("*/textures"),
				shared.JsonValue("*/textures/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTextureId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.TerrainTextureId.References.Get()
			},
		},
	},
}
