package handlers

import (
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var WorldgenTemplatePool = &JsonHandler{
	Pattern: shared.WorldgenTemplatePoolGlob,
	Entries: []JsonEntry{
		{
			Kind:    Definition,
			Binding: stores.WorldgenTemplatePoolId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:template_pool/description/identifier")},
		},
		{
			Kind:    References,
			Binding: stores.WorldgenProcessorId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:template_pool/elements/*/element/processors")},
		},
		{
			Kind:    References,
			Binding: stores.WorldgenTemplatePoolId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:template_pool/fallback")},
		},
		{
			Path: []shared.JsonPath{shared.JsonValue("minecraft:template_pool/elements/*/element/location")},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return sliceutil.Map(stores.StructurePath.Get(), func(s core.Symbol) core.Symbol {
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
