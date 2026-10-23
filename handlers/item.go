package handlers

import (
	"slices"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"golang.org/x/mod/semver"
)

func Item(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/items/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.ItemId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:item/description/identifier")},
				FilterDiff: true,
			},
			{
				Kind:    References,
				Binding: workspace.ItemTextureId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:item/components/minecraft:icon"),
					core.JsonValue("minecraft:item/components/minecraft:icon/texture"),
					core.JsonValue("minecraft:item/components/minecraft:icon/textures/*"),
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.ItemTag,
				Path:    []core.JsonPath{core.JsonValue("minecraft:item/components/minecraft:tags/tags/*")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.ItemTag.Definition.Get(), workspace.ItemTag.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/items/*")},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:item/components/minecraft:block_placer/block"),
					core.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block"),
				},
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.Definition.Get("block")
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.ItemId.References.Get("block")
				},
			},
			{
				Kind:    References,
				Binding: workspace.EntityId,
				Path:    []core.JsonPath{core.JsonValue("minecraft:item/components/minecraft:entity_placer/entity")},
			},
			{
				Kind:    Definition,
				Binding: workspace.CooldownCategory,
				Path:    []core.JsonPath{core.JsonValue("minecraft:item/components/minecraft:cooldown/category")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.CooldownCategory.Definition.Get(), workspace.CooldownCategory.References.Get())
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.ItemCustomComponent,
				Path:    []core.JsonPath{core.JsonValue("minecraft:item/components/minecraft:custom_components/*")},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.ItemCustomComponent.Definition.Get(), workspace.ItemCustomComponent.References.Get())
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.ItemCustomComponent,
				Path:    []core.JsonPath{core.JsonKey("minecraft:item/components/*")},
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
					return slices.Concat(workspace.ItemCustomComponent.Definition.Get(), workspace.ItemCustomComponent.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.Lang,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:item/components/minecraft:display_name"),
					core.JsonValue("minecraft:item/components/minecraft:display_name/value"),
					core.JsonValue("minecraft:item/components/minecraft:interact_button"),
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:item/components/**/condition"),
			core.JsonValue("minecraft:item/components/minecraft:digger/destroy_speeds/*/block/tags"),
			core.JsonValue("minecraft:item/components/minecraft:repairable/repair_items/*/repair_amount"),
			core.JsonValue("minecraft:item/components/minecraft:icon/frame"),
			core.JsonValue("minecraft:item/events/**/sequence/*/condition"),
		},
	}
}
