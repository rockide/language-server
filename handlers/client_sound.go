package handlers

import (
	"github.com/rockide/language-server/core"
)

func ClientSound(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/sounds.json"),
		Entries: []JsonEntry{
			{
				Kind:    References,
				Binding: workspace.SoundDefinitionId,
				Path: []core.JsonPath{
					core.JsonValue("block_sounds/*/events/*"),
					core.JsonValue("block_sounds/*/events/*/sound"),
					core.JsonValue("entity_sounds/entities/*/events/*"),
					core.JsonValue("entity_sounds/entities/*/events/*/sound"),
					core.JsonValue("entity_sounds/entities/*/variants/map/*/events/*"),
					core.JsonValue("entity_sounds/entities/*/variants/map/*/events/*/sound"),
					core.JsonValue("individual_event_sounds/events/*"),
					core.JsonValue("individual_event_sounds/events/*/sound"),
					core.JsonValue("individual_named_sounds/sounds/*"),
					core.JsonValue("individual_named_sounds/sounds/*/sound"),
					core.JsonValue("interactive_sounds/*/*/events/*"),
					core.JsonValue("interactive_sounds/*/*/events/*/sound"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityId,
				Path: []core.JsonPath{
					core.JsonKey("entity_sounds/entities/*"),
					core.JsonKey("interactive_sounds/entity_sounds/entities/*"),
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("entity_sounds/entities/*/variants/key"),
		},
	}
}
