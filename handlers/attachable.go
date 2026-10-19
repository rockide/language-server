package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Attachable = &JsonHandler{
	Pattern: shared.AttachableGlob,
	Entries: []JsonEntry{
		{
			Kind:       References,
			Binding:    stores.ItemId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:       Source,
			Binding:    stores.ClientAnimationAlias,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.ClientAnimationAlias,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
				shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.ClientAnimationId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/materials/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.EntityMaterial,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/materials/*")},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/textures/*")},
			// TODO
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/geometry/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.GeometryId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
		},
		{
			Kind:    References,
			Binding: stores.RenderControllerId,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
				shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/particle_effects/*"),
				shared.JsonKey("minecraft:attachable/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.ParticleId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
				shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/sound_effects/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.SoundDefinitionId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:attachable/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:attachable/description/scripts/scale"),
		shared.JsonValue("minecraft:attachable/description/render_controllers/*/*"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:attachable/description/geometry/*"),
	},
}
