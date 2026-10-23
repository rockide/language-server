package handlers

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"go.lsp.dev/uri"
)

func ClientAnimationController(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/animation_controllers/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.ClientAnimationId,
				Path:       []core.JsonPath{core.JsonKey("animation_controllers/*")},
				FilterDiff: true,
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					filtered := []core.Symbol{}
					for _, ref := range workspace.ClientAnimationId.References.Get() {
						if strings.HasPrefix(ref.Value, "controller.") {
							filtered = append(filtered, ref)
						}
					}
					return filtered
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientAnimationAlias,
				Path: []core.JsonPath{
					core.JsonValue("animation_controllers/*/states/*/animations/*"),
					core.JsonKey("animation_controllers/*/states/*/animations/*/*"),
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
					for _, symbol := range workspace.ClientAnimationId.References.Get(id) {
						if !set.ContainsOne(symbol.URI) {
							set.Add(symbol.URI)
							res = append(res, workspace.ClientAnimationAlias.Definition.GetFrom(symbol.URI)...)
						}
					}
					return res
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					res := workspace.ClientAnimationAlias.References.Get(id)
					set := mapset.NewThreadUnsafeSet[uri.URI]()
					for _, symbol := range workspace.ClientAnimationId.References.Get(id) {
						if !set.ContainsOne(symbol.URI) {
							set.Add(symbol.URI)
							res = append(res, workspace.ClientAnimationAlias.References.GetFrom(symbol.URI)...)
						}
					}
					return res
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.ClientControllerState,
				Path:       []core.JsonPath{core.JsonKey("animation_controllers/*/states/*")},
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
					return workspace.ClientControllerState.Definition.GetFrom(ctx.URI, id)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					return workspace.ClientControllerState.References.GetFrom(ctx.URI, id)
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientControllerState,
				Path: []core.JsonPath{
					core.JsonValue("animation_controllers/*/initial_state"),
					core.JsonKey("animation_controllers/*/states/*/transitions/*/*"),
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
					return workspace.ClientControllerState.Definition.GetFrom(ctx.URI, id)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					return workspace.ClientControllerState.References.GetFrom(ctx.URI, id)
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("animation_controllers/*/states/*/animations/*/*"),
			core.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
			core.JsonValue("animation_controllers/*/states/*/on_entry/*"),
			core.JsonValue("animation_controllers/*/states/*/on_exit/*"),
			core.JsonValue("animation_controllers/*/states/*/parameters/*"),
			core.JsonValue("animation_controllers/*/states/*/particle_effects/*/pre_effect_script"),
			core.JsonValue("animation_controllers/*/states/*/variables/*/input"),
		},
	}
}
