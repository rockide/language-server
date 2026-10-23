package handlers

import (
	"github.com/rockide/language-server/core"
)

func Structure(workspace *core.Workspace) *PathHandler {
	return &PathHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/structures/**/*.mcstructure"),
		Store:     workspace.StructurePath,
	}
}
