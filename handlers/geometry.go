package handlers

import (
	"slices"
	"strings"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
)

func Geometry(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/models/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:    Definition,
				Binding: workspace.GeometryId,
				Path: []core.JsonPath{
					core.JsonKey("*"),
					core.JsonValue("minecraft:geometry/*/description/identifier"),
				},
				Matcher: func(ctx *JsonContext) bool {
					return strings.HasPrefix(ctx.NodeValue, "geometry.")
				},
				Transform: func(value string) string {
					res, _, _ := strings.Cut(value, ":")
					return res
				},
				FilterDiff: true,
			},
			{
				Kind:       Definition,
				Binding:    workspace.GeometryBone,
				Path:       []core.JsonPath{core.JsonValue("minecraft:geometry/*/bones/*/name")},
				FilterDiff: true,
				ScopeKey:   _geometryGetIdentifier,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetIdentifier(ctx)
					return workspace.GeometryBone.Definition.Get(identifier)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetIdentifier(ctx)
					return workspace.GeometryBone.References.Get(identifier)
				},
			},
			{
				// FIXME: Prevent circular references
				Kind:     References,
				Binding:  workspace.GeometryBone,
				Path:     []core.JsonPath{core.JsonValue("minecraft:geometry/*/bones/*/parent")},
				ScopeKey: _geometryGetIdentifier,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetIdentifier(ctx)
					// Prevent self-reference
					parent := ctx.GetParentNode()
					name := jsonc.FindNodeAtLocation(parent, jsonc.Path{"name"})
					result := workspace.GeometryBone.Definition.Get(identifier)
					if name != nil {
						return slices.DeleteFunc(result, func(symbol core.Symbol) bool {
							return symbol.Value == name.Value
						})
					}
					return result
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetIdentifier(ctx)
					return workspace.GeometryBone.References.Get(identifier)
				},
			},
			// Older version of geometry files
			{
				Kind:    Definition,
				Binding: workspace.GeometryBone,
				Path:    []core.JsonPath{core.JsonValue("*/bones/*/name")},
				Matcher: func(ctx *JsonContext) bool {
					return _geometryGetOldIdentifier(ctx) != defaultScope
				},
				FilterDiff: true,
				ScopeKey:   _geometryGetOldIdentifier,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetOldIdentifier(ctx)
					return workspace.GeometryBone.Definition.Get(identifier)
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetOldIdentifier(ctx)
					return workspace.GeometryBone.References.Get(identifier)
				},
			},
			{
				// FIXME: Prevent circular references
				Kind:    References,
				Binding: workspace.GeometryBone,
				Path:    []core.JsonPath{core.JsonValue("*/bones/*/parent")},
				Matcher: func(ctx *JsonContext) bool {
					return _geometryGetOldIdentifier(ctx) != defaultScope
				},
				ScopeKey: _geometryGetOldIdentifier,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetOldIdentifier(ctx)
					// Prevent self-reference
					parent := ctx.GetParentNode()
					name := jsonc.FindNodeAtLocation(parent, jsonc.Path{"name"})
					result := workspace.GeometryBone.Definition.Get(identifier)
					if name != nil {
						return slices.DeleteFunc(result, func(symbol core.Symbol) bool {
							return symbol.Value == name.Value
						})
					}
					return result
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					identifier := _geometryGetOldIdentifier(ctx)
					return workspace.GeometryBone.References.Get(identifier)
				},
			},
		},
		MolangLocations: []core.JsonPath{
			core.JsonValue("minecraft:geometry/*/bones/*/binding"),
		},
		MolangSemanticLocations: []core.JsonPath{
			core.JsonValue("minecraft:geometry/*/description/identifier"),
		},
	}
}

func _geometryGetIdentifier(ctx *JsonContext) string {
	path := ctx.GetPath()
	path = slices.Clone(path[:len(path)-3])
	path = append(path, "description", "identifier")
	node := jsonc.FindNodeAtLocation(ctx.GetRootNode(), jsonc.Path(path))
	if node != nil {
		if id, ok := node.Value.(string); ok {
			return id
		}
	}
	return defaultScope
}

func _geometryGetOldIdentifier(ctx *JsonContext) string {
	node := ctx.GetParentNode().Parent.Parent.Parent.Parent
	identifier, ok := node.Children[0].Value.(string)
	if ok {
		return identifier
	}
	return defaultScope
}
