package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Dialogue = &JsonHandler{
	Pattern: shared.DialogueGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.DialogueId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:npc_dialogue/scenes/*/scene_tag")},
		},
	},
	CommandEntry: JsonCommandEntry{
		Handler:   EmbedDialogueCommand,
		SlashMode: SlashAlways,
		Path: []shared.JsonPath{
			shared.JsonValue("minecraft:npc_dialogue/scenes/*/on_open_commands/*"),
			shared.JsonValue("minecraft:npc_dialogue/scenes/*/on_close_commands/*"),
			shared.JsonValue("minecraft:npc_dialogue/scenes/*/buttons/*/commands/*"),
		},
	},
}
