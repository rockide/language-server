package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var TradeTable = &JsonHandler{
	Pattern:   shared.TradeTableGlob,
	PathStore: stores.TradeTablePath,
	Entries: []JsonEntry{
		{
			Store: stores.ItemId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("tiers/*/groups/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/groups/*/trades/*/wants/*/item"),
				shared.JsonValue("tiers/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/trades/*/wants/*/item"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
	},
}
