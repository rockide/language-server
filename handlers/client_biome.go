package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
)

func ClientBiome(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/biomes/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       References,
				Binding:    workspace.BiomeId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:client_biome/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.SoundDefinitionId,
				Path: sliceutil.Map([]string{
					"minecraft:ambient_sounds/addition",
					"minecraft:ambient_sounds/loop",
					"minecraft:ambient_sounds/mood",
				}, func(value string) core.JsonPath {
					return core.JsonValue("minecraft:client_biome/components/" + value)
				}),
			},
			{
				Kind:    References,
				Binding: workspace.AtmosphereId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:atmosphere_identifier/atmosphere_identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.MusicDefinitionId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:biome_music/music_definition")},
			},
			{
				Kind:    References,
				Binding: workspace.ColorGradingId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:color_grading_identifier/color_grading_identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.FogId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:fog_appearance/fog_identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.LightingId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:lighting_identifier/lighting_identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.WaterId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:client_biome/components/minecraft:water_identifier/water_identifier")},
			},
		},
	}
}
