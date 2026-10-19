package handlers

import (
	"slices"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
	"golang.org/x/mod/semver"
)

var Item = &JsonHandler{
	Pattern: shared.ItemGlob,
	Entries: []JsonEntry{
		{
			Kind:       Definition,
			Binding:    stores.ItemId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:item/description/identifier")},
			FilterDiff: true,
		},
		{
			Kind:    References,
			Binding: stores.ItemTextureId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:icon"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/texture"),
				shared.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
			},
		},
		{
			Kind:    Definition,
			Binding: stores.ItemTag,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:tags/tags/*")},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.ItemTag.Definition.Get(), stores.ItemTag.References.Get())
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:block_placer/block"),
				shared.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Definition.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    References,
			Binding: stores.EntityId,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:entity_placer/entity")},
		},
		{
			Kind:    Definition,
			Binding: stores.CooldownCategory,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:cooldown/category")},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.CooldownCategory.Definition.Get(), stores.CooldownCategory.References.Get())
			},
		},
		{
			Kind:    Definition,
			Binding: stores.ItemCustomComponent,
			Path:    []shared.JsonPath{shared.JsonValue("minecraft:item/components/minecraft:custom_components/*")},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.ItemCustomComponent.Definition.Get(), stores.ItemCustomComponent.References.Get())
			},
		},
		{
			Kind:    Definition,
			Binding: stores.ItemCustomComponent,
			Path:    []shared.JsonPath{shared.JsonKey("minecraft:item/components/*")},
			Matcher: func(ctx *JsonContext) bool {
				root := ctx.GetRootNode()
				formatNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"format_version"})
				if formatNode != nil {
					if version, ok := formatNode.Value.(string); ok {
						match := !strings.HasPrefix(ctx.NodeValue, "minecraft:")
						return semver.Compare("v"+version, "v1.21.80") > 0 && match
					}
				}
				return false
			},
			DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.ItemCustomComponent.Definition.Get(), stores.ItemCustomComponent.References.Get())
			},
		},
		{
			Kind:    References,
			Binding: stores.Lang,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:item/components/minecraft:display_name"),
				shared.JsonValue("minecraft:item/components/minecraft:display_name/value"),
				shared.JsonValue("minecraft:item/components/minecraft:interact_button"),
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:item/components/**/condition"),
		shared.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block/tags"),
		shared.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/repair_amount"),
		shared.JsonValue("minecraft:item/components/minecraft:icon/frame"),
		shared.JsonValue("minecraft:item/events/**/sequence/*/condition"),
	},
}
