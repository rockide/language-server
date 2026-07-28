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
			Store:      stores.ItemId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get()
			},
		},
		{
			Store:      stores.ClientAnimationAlias.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.ClientAnimationAlias.References,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/scripts/animate/*/*"),
				shared.JsonValue("minecraft:attachable/description/scripts/animate/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.ClientAnimationId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/materials/*")},
			// TODO
		},
		{
			Store: stores.EntityMaterial.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/materials/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/textures/*")},
			// TODO
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/textures/*")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/geometry/*")},
			// TODO
		},
		{
			Store: stores.GeometryId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/geometry/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.GeometryId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.GeometryId.References.Get()
			},
		},
		{
			Store: stores.RenderControllerId.References,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:attachable/description/render_controllers/*/*"),
				shared.JsonValue("minecraft:attachable/description/render_controllers/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.References.Get()
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
			Store: stores.ParticleId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:attachable/description/particle_effects/*"),
				shared.JsonValue("minecraft:attachable/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:attachable/description/sound_effects/*")},
			// TODO
		},
		{
			Store: stores.SoundDefinitionId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:attachable/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.References.Get()
			},
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
