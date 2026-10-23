package handlers

import (
	"strings"

	"github.com/rockide/language-server/core"
)

func EntityMaterial(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/materials/entity.material"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.EntityMaterial,
				Path:    []core.JsonPath{core.JsonKey("materials/*")},
				Transform: func(value string) string {
					res, _, _ := strings.Cut(value, ":")
					return res
				},
				FilterDiff: true,
			},
			// TODO: Add base entity materials for inheritance.
		},
	}
}

func ParticleMaterial(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/materials/particles.material"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.ParticleMaterial,
				Path:    []core.JsonPath{core.JsonKey("materials/*")},
				Transform: func(value string) string {
					res, _, _ := strings.Cut(value, ":")
					return res
				},
				FilterDiff: true,
			},
			// TODO: Add base particle materials for inheritance.
		},
	}
}
