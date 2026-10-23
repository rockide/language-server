package handlers

import (
	"github.com/rockide/language-server/core"
)

func Attachable(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/attachables/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       References,
				Binding:    workspace.ItemId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:attachable/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:       Definition,
				Binding:    workspace.ClientAnimationAlias,
				Path:       []core.JsonPath{core.JsonKey("minecraft:attachable/description/animations/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientAnimationAlias,
				Path: []core.JsonPath{
					core.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
					core.JsonValue("minecraft:attachable/description/scripts/animate/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ClientAnimationAlias.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    References,
				Binding: workspace.ClientAnimationId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:attachable/description/animations/*")},
				ScopeKey: func(ctx *JsonContext) string {
					return ctx.NodeValue
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:attachable/description/materials/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.EntityMaterial,
				Path:    []core.JsonPath{core.JsonValue("minecraft:attachable/description/materials/*")},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:attachable/description/textures/*")},
				// TODO
			},
			{
				Path:          []core.JsonPath{core.JsonValue("minecraft:attachable/description/textures/*")},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.TexturePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:attachable/description/geometry/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.GeometryId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:attachable/description/geometry/*")},
			},
			{
				Kind:    References,
				Binding: workspace.RenderControllerId,
				Path: []core.JsonPath{
					core.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
					core.JsonValue("minecraft:attachable/description/render_controllers/*"),
				},
			},
			{
				Path: []core.JsonPath{
					core.JsonKey("minecraft:attachable/description/particle_effects/*"),
					core.JsonKey("minecraft:attachable/description/particle_emitters/*"),
				},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.ParticleId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:attachable/description/particle_effects/*"),
					core.JsonValue("minecraft:attachable/description/particle_emitters/*"),
				},
			},
			{
				Path: []core.JsonPath{core.JsonKey("minecraft:attachable/description/sound_effects/*")},
				// TODO
			},
			{
				Kind:    References,
				Binding: workspace.SoundDefinitionId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:attachable/description/sound_effects/*")},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:attachable/description/scripts/animate/*/*"),
			core.JsonValue("minecraft:attachable/description/scripts/initialize/*"),
			core.JsonValue("minecraft:attachable/description/scripts/parent_setup"),
			core.JsonValue("minecraft:attachable/description/scripts/pre_animation/*"),
			core.JsonValue("minecraft:attachable/description/scripts/scale"),
			core.JsonValue("minecraft:attachable/description/render_controllers/*/*"),
		},
		MolangSemanticLocations: []core.JsonPath{
			core.JsonValue("minecraft:attachable/description/geometry/*"),
		},
	}
}
