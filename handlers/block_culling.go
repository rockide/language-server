package handlers

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/sliceutil"
)

func BlockCulling(workspace *core.Workspace) *JsonHandler {
	return &JsonHandler{
		Workspace: workspace,
		Pattern:   core.ResourcePattern("/block_culling/**/*.json"),
		Entries: []JsonEntry{
			{
				Kind:       Definition,
				Binding:    workspace.BlockCullingId,
				Path:       []core.JsonPath{core.JsonValue("minecraft:block_culling_rules/description/identifier")},
				FilterDiff: true,
			},
			{
				Path:          []core.JsonPath{core.JsonValue("minecraft:block_culling_rules/rules/*/geometry_part/bone")},
				DisableRename: true,
				DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
					geos := _bcGeoIdentifiers(ctx, workspace)
					if geos == nil {
						return nil
					}
					return sliceutil.FlatMap(geos, func(identifier string) []core.Symbol {
						return workspace.GeometryBone.Definition.Get(identifier)
					})
				},
				ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
					geos := _bcGeoIdentifiers(ctx, workspace)
					if geos == nil {
						return nil
					}
					return sliceutil.FlatMap(geos, func(identifier string) []core.Symbol {
						return workspace.GeometryBone.References.Get(identifier)
					})
				},
			},
		},
	}
}

func _bcGeoIdentifiers(ctx *JsonContext, workspace *core.Workspace) []string {
	// Get culilng rule identifier
	root := ctx.GetRootNode()
	identifierNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block_culling_rules", "description", "identifier"})
	if identifierNode == nil {
		return nil
	}
	identifier := identifierNode.Value
	// Find geometries using the culling rules
	var identifiers []string
	set := mapset.NewThreadUnsafeSet[string]()
	for _, ref := range workspace.BlockCullingId.References.Get() {
		document, err := workspace.Documents.GetOrReadFile(ref.URI)
		if err != nil {
			continue
		}
		root, _ := jsonc.ParseTree(document.GetText(), nil)
		var nodes []*jsonc.Node
		// From components
		geometryNode := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block", "components", "minecraft:geometry"})
		if geometryNode != nil {
			nodes = append(nodes, geometryNode)
		}
		// From permutations
		permutations := jsonc.FindNodeAtLocation(root, jsonc.Path{"minecraft:block", "permutations"})
		if permutations != nil && permutations.Type == jsonc.NodeTypeArray {
			for _, child := range permutations.Children {
				geometryNode := jsonc.FindNodeAtLocation(child, jsonc.Path{"components", "minecraft:geometry"})
				if geometryNode != nil {
					nodes = append(nodes, geometryNode)
				}
			}
		}
		// Check each geometry node
		for _, geometryNode := range nodes {
			identifierNode := jsonc.FindNodeAtLocation(geometryNode, jsonc.Path{"identifier"})
			if identifierNode == nil {
				continue
			}
			cullingRulesNode := jsonc.FindNodeAtLocation(geometryNode, jsonc.Path{"culling"})
			if cullingRulesNode == nil {
				continue
			}
			cullingIdentifier, ok := cullingRulesNode.Value.(string)
			if !ok || cullingIdentifier != identifier {
				continue
			}
			geometryIdentifier, ok := identifierNode.Value.(string)
			if !ok {
				continue
			}
			if !set.Contains(geometryIdentifier) {
				set.Add(geometryIdentifier)
				identifiers = append(identifiers, geometryIdentifier)
			}
		}
	}
	return identifiers
}
