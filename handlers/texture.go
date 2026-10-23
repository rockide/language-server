package handlers

import (
	"github.com/rockide/language-server/core"
)

func Texture(workspace *core.Workspace) *PathHandler {
	return &PathHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/textures/**/*.{png,tga,jpg,jpeg}"),
		Store:     workspace.TexturePath,
	}
}
