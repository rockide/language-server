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
			Store:      stores.EntityId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityId.References.Get()
			},
		},
		{
			Store:      stores.ClientAnimationAlias.Source,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/animations/*")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
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
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
		},
		{
			Store: stores.ClientAnimationAlias.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/scripts/animate/*"),
				shared.JsonKey("minecraft:client_entity/description/scripts/animate/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ClientAnimationAlias.Source.GetFrom(ctx.URI)
			},
			References: func(ctx *JsonContext) []core.Symbol {
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
			Store: stores.ClientAnimationId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/animations/*")},
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
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/materials/*")},
			// TODO
		},
		{
			Store: stores.EntityMaterial.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/materials/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.EntityMaterial.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/textures/*")},
		},
		{
			Path:          []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/textures/*")},
			DisableRename: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.TexturePath.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/geometry/*")},
			// TODO
		},
		{
			Store: stores.GeometryId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/geometry/*")},
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
				shared.JsonValue("minecraft:client_entity/description/render_controllers/*"),
				shared.JsonKey("minecraft:client_entity/description/render_controllers/*/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.RenderControllerId.References.Get()
			},
		},
		{
			Store: stores.ItemTextureId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/spawn_egg/texture")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTextureId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemTextureId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonKey("minecraft:client_entity/description/particle_emitters/*"),
			},
			// TODO
		},
		{
			Store: stores.ParticleId.References,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:client_entity/description/particle_effects/*"),
				shared.JsonValue("minecraft:client_entity/description/particle_emitters/*"),
			},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ParticleId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonKey("minecraft:client_entity/description/sound_effects/*")},
			// TODO
		},
		{
			Store: stores.SoundDefinitionId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_entity/description/sound_effects/*")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.References.Get()
			},
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
