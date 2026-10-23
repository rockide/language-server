package handlers

import (
	"github.com/rockide/language-server/core"
)

func Sound(workspace *core.Workspace) *PathHandler {
	return &PathHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/sounds/**/*.{fsb,ogg,wav}"),
		Store:     workspace.SoundPath,
	}
}
