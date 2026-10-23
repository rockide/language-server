package handlers

import (
	"github.com/rockide/language-server/core"
)

func Dialogue(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/dialogue/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.DialogueId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:npc_dialogue/scenes/*/scene_tag")},
			},
		},
		CommandEntry: JsonCommandEntry{
			Handler:   EmbedDialogueCommand(workspace),
			SlashMode: SlashAlways,
			Path: []core.JsonPath{
				core.JsonValue("minecraft:npc_dialogue/scenes/*/on_open_commands/*"),
				core.JsonValue("minecraft:npc_dialogue/scenes/*/on_close_commands/*"),
				core.JsonValue("minecraft:npc_dialogue/scenes/*/buttons/*/commands/*"),
			},
		},
	}
}
