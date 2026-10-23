package handlers

import (
	"github.com/rockide/language-server/core"
)

func TradeTable(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/trading/**/*.json"),
		PathStore: workspace.TradeTablePath,
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("tiers/*/groups/*/trades/*/gives/*/item"),
					core.JsonValue("tiers/*/groups/*/trades/*/wants/*/item"),
					core.JsonValue("tiers/*/trades/*/gives/*/item"),
					core.JsonValue("tiers/*/trades/*/wants/*/item"),
				},
			},
		},
	}
}
