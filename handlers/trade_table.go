package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var TradeTable = &JsonHandler{
	Pattern:   shared.TradeTableGlob,
	PathStore: stores.TradeTablePath,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("tiers/*/groups/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/groups/*/trades/*/wants/*/item"),
				shared.JsonValue("tiers/*/trades/*/gives/*/item"),
				shared.JsonValue("tiers/*/trades/*/wants/*/item"),
			},
		},
	},
}
