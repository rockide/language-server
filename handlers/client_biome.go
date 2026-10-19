package handlers

import (
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientBiome = &JsonHandler{
	Pattern: shared.ClientBiomeGlob,
	Entries: []JsonEntry{
		{
			Kind:       References,
			Binding:    stores.BiomeId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:client_biome/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.SoundDefinitionId,
			Path: sliceutil.Map([]string{
				"minecraft:ambient_sounds/addition",
				"minecraft:ambient_sounds/loop",
				"minecraft:ambient_sounds/mood",
			}, func(value string) shared.JsonPath {
				return shared.JsonValue("minecraft:client_biome/components/" + value)
			}),
		},
		{
			Kind:    References,
			Binding: stores.AtmosphereId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:atmosphere_identifier/atmosphere_identifier")},
		},
		{
			Kind:    References,
			Binding: stores.MusicDefinitionId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:biome_music/music_definition")},
		},
		{
			Kind:    References,
			Binding: stores.ColorGradingId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:color_grading_identifier/color_grading_identifier")},
		},
		{
			Kind:    References,
			Binding: stores.FogId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:fog_appearance/fog_identifier")},
		},
		{
			Kind:    References,
			Binding: stores.LightingId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:lighting_identifier/lighting_identifier")},
		},
		{
			Kind:    References,
			Binding: stores.WaterId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:client_biome/components/minecraft:water_identifier/water_identifier")},
		},
	},
}
