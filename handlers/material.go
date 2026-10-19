package handlers

import (
	"strings"

	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var EntityMaterial = &JsonHandler{
	Pattern: shared.EntityMaterialGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.EntityMaterial,
			Path:    []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
		},
		// TODO: Add base entity materials for inheritance.
	},
}

var ParticleMaterial = &JsonHandler{
	Pattern: shared.ParticleMaterialGlob,
	Entries: []JsonEntry{
		{
			Kind:    Source,
			Binding: stores.ParticleMaterial,
			Path:    []shared.JsonPath{shared.JsonKey("materials/*")},
			Transform: func(value string) string {
				res, _, _ := strings.Cut(value, ":")
				return res
			},
			FilterDiff: true,
		},
		// TODO: Add base particle materials for inheritance.
	},
}
