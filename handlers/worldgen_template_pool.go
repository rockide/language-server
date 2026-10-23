package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
)

func WorldgenTemplatePool(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/worldgen/template_pools/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.WorldgenTemplatePoolId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:template_pool/description/identifier")},
			},
			{
				Kind:    References,
				Binding: workspace.WorldgenProcessorId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:template_pool/elements/*/element/processors")},
			},
			{
				Kind:    References,
				Binding: workspace.WorldgenTemplatePoolId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:template_pool/fallback")},
			},
			{
				Path: []core.JsonPath{core.JsonValue("minecraft:template_pool/elements/*/element/location")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return sliceutil.Map(workspace.StructurePath.Get(), func(s core.Symbol) core.Symbol {
						s.Value = s.Value[11:]
						return s
					})
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
		},
	}
}
