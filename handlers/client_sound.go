package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var ClientSound = &JsonHandler{
	Pattern: shared.ClientSoundGlob,
	Entries: []JsonEntry{
		{
			Kind:    References,
			Binding: stores.SoundDefinitionId,
			Path: []shared.JsonPath{
				shared.JsonValue("block_sounds/*/events/*"),
				shared.JsonValue("block_sounds/*/events/*/sound"),
				shared.JsonValue("entity_sounds/entities/*/events/*"),
				shared.JsonValue("entity_sounds/entities/*/events/*/sound"),
				shared.JsonValue("entity_sounds/entities/*/variants/map/*/events/*"),
				shared.JsonValue("entity_sounds/entities/*/variants/map/*/events/*/sound"),
				shared.JsonValue("individual_event_sounds/events/*"),
				shared.JsonValue("individual_event_sounds/events/*/sound"),
				shared.JsonValue("individual_named_sounds/sounds/*"),
				shared.JsonValue("individual_named_sounds/sounds/*/sound"),
				shared.JsonValue("interactive_sounds/*/*/events/*"),
				shared.JsonValue("interactive_sounds/*/*/events/*/sound"),
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityId,
			Path: []shared.JsonPath{
				shared.JsonKey("entity_sounds/entities/*"),
				shared.JsonKey("interactive_sounds/entity_sounds/entities/*"),
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("entity_sounds/entities/*/variants/key"),
	},
}
