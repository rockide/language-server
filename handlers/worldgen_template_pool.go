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
			Store: stores.WorldgenTemplatePoolId.Source,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/description/identifier")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.References.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.Source.Get()
			},
		},
		{
			Store: stores.WorldgenProcessorId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/elements/*/element/processors")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessorId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenProcessorId.References.Get()
			},
		},
		{
			Store: stores.WorldgenTemplatePoolId.References,
			Path:  []shared.JsonPath{shared.JsonValue("minecraft:template_pool/fallback")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.Source.Get()
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return stores.WorldgenTemplatePoolId.References.Get()
			},
		},
		{
			Path: []shared.JsonPath{shared.JsonValue("minecraft:template_pool/elements/*/element/location")},
			Source: func(ctx *JsonContext) []core.Symbol {
				return sliceutil.Map(stores.StructurePath.Get(), func(s core.Symbol) core.Symbol {
					s.Value = s.Value[11:]
					return s
				})
			},
			References: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
	},
}
