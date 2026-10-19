package handlers

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientEntity = &JsonHandler{
	Pattern: shared.ClientEntityGlob,
	Entries: []JsonEntry{
		{
			Kind:       References,
			Binding:    stores.EntityId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:       Source,
			Binding:    stores.ClientAnimationAlias,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				res := stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
				set := mapset.NewThreadUnsafeSet[string]()
				for _, symbol := range stores.ClientAnimationId.References.GetFrom(ctx.URI) {
					if !set.ContainsOne(symbol.Value) {
						set.Add(symbol.Value)
						res = append(res, stores.ClientAnimationAlias.References.Get(symbol.Value)...)
					}
				}
				return res
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    References,
			Binding: stores.ClientAnimationAlias,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				res := stores.ClientAnimationAlias.References.GetFrom(ctx.URI)
				set := mapset.NewThreadUnsafeSet[string]()
				for _, symbol := range stores.ClientAnimationId.References.GetFrom(ctx.URI) {
					if !set.ContainsOne(symbol.Value) {
						set.Add(symbol.Value)
						res = append(res, stores.ClientAnimationAlias.References.Get(symbol.Value)...)
					}
				}
				return res
			},
		},
		{
			Kind:    References,
			Binding: stores.ClientAnimationId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
			ScopeKey: func(ctx *JsonContext) string {
				return ctx.NodeValue
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/materials/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.EntityMaterial,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/materials/*")},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/textures/*")},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/geometry/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.GeometryId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
		},
		{
			Kind:    References,
			Binding: stores.RenderControllerId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
				shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemTextureId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.ParticleId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/sound_effects/*")},
			// TODO
		},
		{
			Kind:    References,
			Binding: stores.SoundDefinitionId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/scripts/animate/*/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/initialize/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/parent_setup"),
		shared.JsonValue("minecraft:client_entity/description/scripts/pre_animation/*"),
		shared.JsonValue("minecraft:client_entity/description/scripts/scale"),
		shared.JsonValue("minecraft:client_entity/description/render_controllers/*/*"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:client_entity/description/geometry/*"),
	},
}
