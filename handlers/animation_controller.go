package handlers

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
	"go.lsp.dev/uri"
)

var AnimationController = &JsonHandler{
	Pattern: shared.AnimationControllerGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.AnimationId,
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*")},
			FilterDiff: true,
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				filtered := []core.Symbol{}
				for _, ref := range stores.AnimationId.References.Get() {
					if strings.HasPrefix(ref.Value, "controller.") {
						filtered = append(filtered, ref)
					}
				}
				return filtered
			},
		},
		{
			Kind:    References,
			Binding: stores.AnimationAlias,
			Path: []shared.JsonPath{
				shared.JsonValue("animation_controllers/*/states/*/animations/*"),
				shared.JsonKey("animation_controllers/*/states/*/animations/*/*"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[1].(string); ok {
					return id
				}
				return defaultScope
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				res := []core.Symbol{}
				set := mapset.NewThreadUnsafeSet[uri.URI]()
				for _, symbol := range stores.AnimationId.References.Get(id) {
					if !set.ContainsOne(symbol.URI) {
						set.Add(symbol.URI)
						res = append(res, stores.AnimationAlias.Definition.GetFrom(symbol.URI)...)
					}
				}
				return res
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				res := stores.AnimationAlias.References.Get(id)
				set := mapset.NewThreadUnsafeSet[uri.URI]()
				for _, symbol := range stores.AnimationId.References.Get(id) {
					if !set.ContainsOne(symbol.URI) {
						set.Add(symbol.URI)
						res = append(res, stores.AnimationAlias.References.GetFrom(symbol.URI)...)
					}
				}
				return res
			},
		},
		{
			Kind:       Definition,
			Binding:    stores.ControllerState,
			Path:       []shared.JsonPath{shared.JsonKey("animation_controllers/*/states/*")},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[1].(string); ok {
					return id
				}
				return defaultScope
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ControllerState.Definition.GetFrom(ctx.URI, id)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ControllerState.References.GetFrom(ctx.URI, id)
			},
		},
		{
			Kind:    References,
			Binding: stores.ControllerState,
			Path: []shared.JsonPath{
				shared.JsonValue("animation_controllers/*/initial_state"),
				shared.JsonKey("animation_controllers/*/states/*/transitions/*/*"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				if id, ok := ctx.GetPath()[1].(string); ok {
					return id
				}
				return defaultScope
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ControllerState.Definition.GetFrom(ctx.URI, id)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				id, ok := ctx.GetPath()[1].(string)
				if !ok {
					return nil
				}
				return stores.ControllerState.References.GetFrom(ctx.URI, id)
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("animation_controllers/*/states/*/animations/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_entry/*"),
		shared.JsonValue("animation_controllers/*/states/*/on_exit/*"),
	},
	CommandEntry: JsonCommandEntry{
		Handler:   EmbedEventCommand,
		SlashMode: SlashAlways,
		Path: []shared.JsonPath{
			shared.JsonValue("animation_controllers/*/states/*/animations/*/*"),
			shared.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
			shared.JsonValue("animation_controllers/*/states/*/on_entry/*"),
			shared.JsonValue("animation_controllers/*/states/*/on_exit/*"),
		},
	},
}
