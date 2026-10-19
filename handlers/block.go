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

var Block = &JsonHandler{
	Pattern: shared.BlockGlob,
	Entries: []JsonEntry{
		{
			Kind:       Source,
			Binding:    stores.ItemId,
			Path:       []shared.JsonPath{shared.JsonValue("minecraft:block/description/identifier")},
			FilterDiff: true,
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
		},
		{
			Kind:       Source,
			Binding:    stores.BlockState,
			Path:       []shared.JsonPath{shared.JsonKey("minecraft:block/description/states/*")},
			FilterDiff: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockState.References.GetFrom(ctx.URI)
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.BlockState.Source.GetFrom(ctx.URI)
			},
		},
		{
			Kind:    Source,
			Binding: stores.BlockTag,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:tags/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:tags/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockTag.Source.Get(), stores.BlockTag.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    Source,
			Binding: stores.BlockTag,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:block/components/*"),
				shared.JsonKey("minecraft:block/permutations/*/components/*"),
			},
			Matcher: func(ctx *JsonContext) bool {
				return strings.HasPrefix(ctx.NodeValue, "tag:")
			},
			Transform: func(value string) string {
				res, _ := strings.CutPrefix(value, "tag:")
				return res
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockTag.Source.Get(), stores.BlockTag.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.GeometryId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
			},
		},
		{
			Kind:    References,
			Binding: stores.TerrainTextureId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:material_instances/*/texture"),
				shared.JsonValue("minecraft:block/components/minecraft:item_visual/material_instances/*/texture"),
				shared.JsonValue("minecraft:block/components/minecraft:destruction_particles/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/material_instances/*/texture"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:destruction_particles/texture"),
			},
		},
		{
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:loot"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:loot"),
			},
			DisableRename: true,
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.LootTablePath.Get()
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.BlockCullingId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:geometry/culling"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/culling"),
			},
		},
		{
			Kind:    Source,
			Binding: stores.RecipeTag,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:crafting_table/crafting_tags/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:crafting_table/crafting_tags/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.RecipeTag.Source.Get(), stores.RecipeTag.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.ItemId,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/block_filter/*/name"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*"),
				shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/block_filter/*/name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/name"),
			},
			ScopeKey: func(ctx *JsonContext) string {
				return "block"
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.Source.Get("block")
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return stores.ItemId.References.Get("block")
			},
		},
		{
			Kind:    Source,
			Binding: stores.BlockCustomComponent,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:custom_components/*"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:custom_components/*"),
			},
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockCustomComponent.Source.Get(), stores.BlockCustomComponent.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    Source,
			Binding: stores.BlockCustomComponent,
			Path: []shared.JsonPath{
				shared.JsonKey("minecraft:block/components/*"),
				shared.JsonKey("minecraft:block/permutations/*/components/*"),
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
			SourceOverride: func(ctx *JsonContext) []core.Symbol {
				return slices.Concat(stores.BlockCustomComponent.Source.Get(), stores.BlockCustomComponent.References.Get())
			},
			ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
				return nil
			},
		},
		{
			Kind:    References,
			Binding: stores.Lang,
			Path: []shared.JsonPath{
				shared.JsonValue("minecraft:block/components/minecraft:display_name"),
				shared.JsonValue("minecraft:block/components/minecraft:display_name/value"),
				shared.JsonValue("minecraft:block/components/minecraft:crafting_table/table_name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:display_name"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:display_name/value"),
				shared.JsonValue("minecraft:block/permutations/*/components/minecraft:crafting_table/table_name"),
			},
		},
	},
	MolangLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:destructible_by_mining/item_specific_speeds/*/item/tags"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/bone_visibility/*"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:placement_filter/conditions/*/block_filter/*/tags"),
		shared.JsonValue("minecraft:block/permutations/*/condition"),
	},
	MolangSemanticLocations: []shared.JsonPath{
		shared.JsonValue("minecraft:block/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/components/minecraft:geometry/identifier"),
		shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry"),
		shared.JsonValue("minecraft:block/components/minecraft:item_visual/geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:geometry/identifier"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry"),
		shared.JsonValue("minecraft:block/permutations/*/components/minecraft:item_visual/geometry/identifier"),
	},
}
