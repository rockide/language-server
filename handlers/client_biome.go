package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientBiome = &JsonHandler{
	Pattern: shared.ClientBiomeGlob,
	Entries: []JsonEntry{
		{
			Store:      stores.BiomeId.References,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_biome/description/identifier")},
			FilterDiff: true,
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.BiomeId.References.Get()
			},
		},
		{
			Store: stores.SoundDefinitionId.References,
			Path: sliceutil.Map([]string{
				"minecraft:ambient_sounds/addition",
				"minecraft:ambient_sounds/loop",
				"minecraft:ambient_sounds/mood",
			}, func(value string) shared.JsonPath {
				return shared.JsonValue("minecraft:client_biome/components/" + value)
			}),
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.SoundDefinitionId.References.Get()
			},
		},
		{
			Store: stores.AtmosphereId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:atmosphere_identifier/atmosphere_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.AtmosphereId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.AtmosphereId.References.Get()
			},
		},
		{
			Store: stores.MusicDefinitionId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:biome_music/music_definition")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinitionId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.MusicDefinitionId.References.Get()
			},
		},
		{
			Store: stores.ColorGradingId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:color_grading_identifier/color_grading_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGradingId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.ColorGradingId.References.Get()
			},
		},
		{
			Store: stores.FogId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:fog_appearance/fog_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.FogId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.FogId.References.Get()
			},
		},
		{
			Store: stores.LightingId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:lighting_identifier/lighting_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.LightingId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.LightingId.References.Get()
			},
		},
		{
			Store: stores.WaterId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:water_identifier/water_identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WaterId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WaterId.References.Get()
			},
		},
	},
}
