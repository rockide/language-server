package handlers

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
)

func ClientEntity(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/entity/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       References,
				Binding:    workspace.EntityId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:client_entity/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:       Definition,
				Binding:    workspace.ClientAnimationAlias,
				Path:       []core.JsonPath{core.JsonKey("minecraft:client_entity/description/animations/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					res := workspace.ClientAnimationAlias.References.GetFrom(ctx.URI)
					set := mapset.NewThreadUnsafeSet[string]()
					for _, symbol := range workspace.ClientAnimationId.References.GetFrom(ctx.URI) {
						if !set.ContainsOne(symbol.Value) {
							set.Add(symbol.Value)
							res = append(res, workspace.ClientAnimationAlias.References.Get(symbol.Value)...)
						}
					}
					return res
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientAnimationAlias,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
					core.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					res := workspace.ClientAnimationAlias.References.GetFrom(ctx.URI)
					set := mapset.NewThreadUnsafeSet[string]()
					for _, symbol := range workspace.ClientAnimationId.References.GetFrom(ctx.URI) {
						if !set.ContainsOne(symbol.Value) {
							set.Add(symbol.Value)
							res = append(res, workspace.ClientAnimationAlias.References.Get(symbol.Value)...)
						}
					}
					return res
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientAnimationId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_entity/description/animations/*")},
				ScopeKey: func(ctx *JsonContext) string {
					return ctx.NodeValue
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:client_entity/description/materials/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.EntityMaterial,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_entity/description/materials/*")},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:client_entity/description/textures/*")},
			},
			{
				Path:          []core.JsonPath{core.JsonValue("minecraft:client_entity/description/textures/*")},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.TexturePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:client_entity/description/geometry/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.GeometryId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_entity/description/geometry/*")},
			},
			{
				Kind:    References,
				Binding: workspace.RenderControllerId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:client_entity/description/render_controllers/*"),
					core.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemTextureId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
			},
			{
				Path: []core.JsonPath{
					core.JsonKey("minecraft:client_entity/description/particle_effects/*"),
					core.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
				},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.ParticleId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:client_entity/description/particle_effects/*"),
					core.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:client_entity/description/sound_effects/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.SoundDefinitionId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_entity/description/sound_effects/*")},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:client_entity/description/scripts/animate/*/*"),
			core.JsonValue("minecraft:client_entity/description/scripts/initialize/*"),
			core.JsonValue("minecraft:client_entity/description/scripts/parent_setup"),
			core.JsonValue("minecraft:client_entity/description/scripts/pre_animation/*"),
			core.JsonValue("minecraft:client_entity/description/scripts/scale"),
			core.JsonValue("minecraft:client_entity/description/render_controllers/*/*"),
		},
		MolangSemanticLocations: []core.JsonPath{
			core.JsonValue("minecraft:client_entity/description/geometry/*"),
		},
	}
}
