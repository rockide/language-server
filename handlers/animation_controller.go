package handlers

import (
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"go.lsp.dev/uri"
)

func AnimationController(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/animation_controllers/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.AnimationId,
				Path:       []core.JsonPath{core.JsonKey("animation_controllers/*")},
				FilterDiff: true,
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					filtered := []core.Symbol{}
					for _, ref := range workspace.AnimationId.References.Get() {
						if strings.HasPrefix(ref.Value, "controller.") {
							filtered = append(filtered, ref)
						}
					}
					return filtered
				},
			},
			{
				Kind:    References,
				Binding: workspace.AnimationAlias,
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
					for _, symbol := range workspace.AnimationId.References.Get(id) {
						if !set.ContainsOne(symbol.URI) {
							set.Add(symbol.URI)
							res = append(res, workspace.AnimationAlias.Definition.GetFrom(symbol.URI)...)
						}
					}
					return res
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					res := workspace.AnimationAlias.References.Get(id)
					set := mapset.NewThreadUnsafeSet[uri.URI]()
					for _, symbol := range workspace.AnimationId.References.Get(id) {
						if !set.ContainsOne(symbol.URI) {
							set.Add(symbol.URI)
							res = append(res, workspace.AnimationAlias.References.GetFrom(symbol.URI)...)
						}
					}
					return res
				},
			},
			{
				Kind:       Definition,
				Binding:    workspace.ControllerState,
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
					return workspace.ControllerState.Definition.GetFrom(ctx.URI, id)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					return workspace.ControllerState.References.GetFrom(ctx.URI, id)
				},
			},
			{
				Kind:    References,
				Binding: workspace.ControllerState,
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
					return workspace.ControllerState.Definition.GetFrom(ctx.URI, id)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					id, ok := ctx.GetPath()[1].(string)
					if !ok {
						return nil
					}
					return workspace.ControllerState.References.GetFrom(ctx.URI, id)
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("animation_controllers/*/states/*/animations/*/*"),
			core.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
			core.JsonValue("animation_controllers/*/states/*/on_entry/*"),
			core.JsonValue("animation_controllers/*/states/*/on_exit/*"),
		},
		CommandEntry: JsonCommandEntry{
			Handler:   EmbedEventCommand(workspace),
			SlashMode: SlashAlways,
			Path: []core.JsonPath{
				core.JsonValue("animation_controllers/*/states/*/animations/*/*"),
				core.JsonValue("animation_controllers/*/states/*/transitions/*/*"),
				core.JsonValue("animation_controllers/*/states/*/on_entry/*"),
				core.JsonValue("animation_controllers/*/states/*/on_exit/*"),
			},
		},
	}
}
