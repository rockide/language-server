package handlers

import (
	"github.com/rockide/language-server/core"
)

func WorldgenJigsaw(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/worldgen/structures/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.WorldgenJigsawId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:jigsaw/description/identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.BiomeTag,
				Path:    []core.JsonPath{core.JsonValue("minecraft:jigsaw/biome_filters/**/value")},
			},
			{
				Kind:    References,
				Binding: workspace.WorldgenTemplatePoolId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:jigsaw/start_pool"),
					core.JsonValue("minecraft:jigsaw/pool_aliases/*/alias"),
					core.JsonValue("minecraft:jigsaw/pool_aliases/*/target"),
					core.JsonValue("minecraft:jigsaw/pool_aliases/*/targets/*/data"),
					core.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/alias"),
					core.JsonValue("minecraft:jigsaw/pool_aliases/*/groups/*/data/*/target"),
				},
			},
		},
	}
}
