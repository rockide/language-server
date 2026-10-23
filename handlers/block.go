package handlers

import (
	"slices"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"golang.org/x/mod/semver"
)

func Block(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.BehaviorPattern("/blocks/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.ItemId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:block/description/identifier")},
				FilterDiff: true,
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
				Kind:       Definition,
				Binding:    workspace.BlockState,
				Path:       []core.JsonPath{core.JsonKey("minecraft:block/description/states/*")},
				FilterDiff: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.BlockState.Definition.GetFrom(ctx.URI)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.BlockState.References.GetFrom(ctx.URI)
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.BlockTag,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:tags/*"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:tags/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.BlockTag.Definition.Get(), workspace.BlockTag.References.Get())
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.BlockTag,
				Path: []core.JsonPath{
					core.JsonKey("minecraft:block/components/*"),
					core.JsonKey("minecraft:block/permutations/*/components/*"),
				},
				Matcher: func(ctx *JsonContext) bool {
					return strings.HasPrefix(ctx.NodeValue, "tag:")
				},
				Transform: func(value string) string {
					res, _ := strings.CutPrefix(value, "tag:")
					return res
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.BlockTag.Definition.Get(), workspace.BlockTag.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.GeometryId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:geometry"),
					core.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
					core.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
					core.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
				},
			},
			{
				Kind:    References,
				Binding: workspace.TerrainTextureId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:material_instances/*/texture"),
					core.JsonValue("minecraft:block/components/minecraft:item_visual/material_instances/*/texture"),
					core.JsonValue("minecraft:block/components/minecraft:destruction_particles/texture"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:material_instances/*/texture"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/material_instances/*/texture"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:destruction_particles/texture"),
				},
			},
			{
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:loot"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:loot"),
				},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return workspace.LootTablePath.Get()
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
			},
			{
				Kind:    References,
				Binding: workspace.BlockCullingId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:geometry/culling"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/culling"),
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.RecipeTag,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:crafting_table/crafting_tags/*"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:crafting_table/crafting_tags/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.RecipeTag.Definition.Get(), workspace.RecipeTag.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.ItemId,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*"),
					core.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*/name"),
					core.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*"),
					core.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*/name"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
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
				Kind:    Definition,
				Binding: workspace.BlockCustomComponent,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:custom_components/*"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:custom_components/*"),
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.BlockCustomComponent.Definition.Get(), workspace.BlockCustomComponent.References.Get())
				},
			},
			{
				Kind:    Definition,
				Binding: workspace.BlockCustomComponent,
				Path: []core.JsonPath{
					core.JsonKey("minecraft:block/components/*"),
					core.JsonKey("minecraft:block/permutations/*/components/*"),
				},
				Matcher: func(ctx *JsonContext) bool {
					root := ctx.GetRootNode()
					formatNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"format_version"})
					if formatNode != nil {
						if version, ok := formatNode.Value.(string); ok {
							match := !strings.HasPrefix(ctx.NodeValue, "minecraft:") && !strings.HasPrefix(ctx.NodeValue, "tag:")
							return semver.Compare("v"+version, "v1.21.80") > 0 && match
						}
					}
					return false
				},
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					return nil
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					return slices.Concat(workspace.BlockCustomComponent.Definition.Get(), workspace.BlockCustomComponent.References.Get())
				},
			},
			{
				Kind:    References,
				Binding: workspace.Lang,
				Path: []core.JsonPath{
					core.JsonValue("minecraft:block/components/minecraft:display_name"),
					core.JsonValue("minecraft:block/components/minecraft:display_name/value"),
					core.JsonValue("minecraft:block/components/minecraft:crafting_table/table_name"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:display_name"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:display_name/value"),
					core.JsonValue("minecraft:block/permutations/*/components/minecraft:crafting_table/table_name"),
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:block/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
			core.JsonValue("minecraft:block/components/minecraft:geometry/bone_visibility/*"),
			core.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/bone_visibility/*"),
			core.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/bone_visibility/*"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
			core.JsonValue("minecraft:block/permutations/*/condition"),
		},
		MolangSemanticLocations: []core.JsonPath{
			core.JsonValue("minecraft:block/components/minecraft:geometry"),
			core.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
			core.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
			core.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
			core.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
		},
	}
}
