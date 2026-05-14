package handlers

import (
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/jsonc"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/protocol/semtok"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

const defaultScope = "__default__"

type JsonContext struct {
	URI           protocol.DocumentURI
	NodeValue     string
	GetPath       func() jsonc.Path
	GetParentNode func() *jsonc.Node
	GetRootNode   func() *jsonc.Node
}

type JsonEntry struct {
	Store         *stores.SymbolStore
	Path          []shared.JsonPath
	Matcher       func(ctx *JsonContext) bool
	Transform     func(value string) string
	DisableRename bool
	// Filter completions to only show undeclared reference
	FilterDiff bool
	// Function to set the scope key. If omitted will use the `defaultScope` value
	ScopeKey func(ctx *JsonContext) string
	// Source for completions and definitions
	Source func(ctx *JsonContext) []core.Symbol
	// References that uses the same source
	References func(ctx *JsonContext) []core.Symbol
}

type JsonHandler struct {
	Pattern                 shared.Pattern
	PathStore               *stores.PathStore
	Entries                 []JsonEntry
	MolangLocations         []shared.JsonPath
	MolangSemanticLocations []shared.JsonPath
	CommandEntry            JsonCommandEntry
}

type JsonCommandEntry struct {
	Path         []shared.JsonPath
	RequireSlash bool
	Handler      *CommandHandler
}

func (j *JsonHandler) GetPattern() shared.Pattern {
	return j.Pattern
}

func (j *JsonHandler) Parse(uri protocol.DocumentURI) error {
	if j.PathStore != nil {
		j.PathStore.Insert(j.Pattern, uri)
	}
	document, err := textdocument.GetOrReadFile(uri)
	if err != nil {
		return err
	}
	return j.ParseDocument(document)
}

func (j *JsonHandler) ParseDocument(document *textdocument.TextDocument) error {
	root, _ := jsonc.ParseTree(document.GetText(), nil)

	commandRanges := []protocol.Range{}
	for _, entry := range j.CommandEntry.Path {
		for _, node := range entry.GetNodes(root) {
			nodeValue, ok := node.Value.(string)
			if !ok {
				continue
			}
			if j.CommandEntry.RequireSlash {
				if len(nodeValue) == 0 || nodeValue[0] != '/' {
					continue
				}
			}
			commandRanges = append(commandRanges, protocol.Range{
				Start: document.PositionAt(node.Offset + 1),
				End:   document.PositionAt(node.Offset + node.Length - 1),
			})
		}
	}
	if len(commandRanges) > 0 {
		commandDocument := document.CreateVirtualDocument(commandRanges...)
		j.CommandEntry.Handler.ParseDocument(commandDocument)
	}

	for _, entry := range j.Entries {
		if entry.Store == nil {
			continue
		}
		for _, jsonPath := range entry.Path {
			for _, node := range jsonPath.GetNodes(root) {
				nodeValue, ok := node.Value.(string)
				if !ok {
					continue
				}
				ctx := JsonContext{
					URI:       document.URI,
					NodeValue: nodeValue,
					GetPath: func() jsonc.Path {
						return jsonc.GetNodePath(node)
					},
					GetParentNode: func() *jsonc.Node {
						path := jsonc.GetNodePath(node)
						return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
					},
					GetRootNode: func() *jsonc.Node {
						return root
					},
				}
				if entry.Matcher != nil && !entry.Matcher(&ctx) {
					continue
				}
				if entry.Transform != nil {
					nodeValue = entry.Transform(nodeValue)
				}
				scope := defaultScope
				if entry.ScopeKey != nil {
					scope = entry.ScopeKey(&ctx)
				}
				entry.Store.Insert(scope, core.Symbol{
					Value: nodeValue,
					URI:   document.URI,
					Range: &protocol.Range{
						Start: document.PositionAt(node.Offset + 1),
						End:   document.PositionAt(node.Offset + node.Length - 1),
					},
				})
			}
		}
	}
	return nil
}

func (j *JsonHandler) Delete(uri protocol.DocumentURI) {
	if j.CommandEntry.Handler != nil {
		j.CommandEntry.Handler.Delete(uri)
	}
	for _, entry := range j.Entries {
		if entry.Store != nil {
			entry.Store.Delete(uri)
		}
	}
}

func (j *JsonHandler) prepareContext(document *textdocument.TextDocument, location *jsonc.Location) (*JsonEntry, *JsonContext) {
	var nodeValue string
	if node := location.PreviousNode; node != nil {
		nodeValue, _ = node.Value.(string)
	}
	params := JsonContext{
		URI:       document.URI,
		NodeValue: nodeValue,
		GetPath: func() jsonc.Path {
			return location.Path
		},
		GetParentNode: func() *jsonc.Node {
			root, _ := jsonc.ParseTree(document.GetText(), nil)
			path := location.Path
			return jsonc.FindNodeAtLocation(root, path[:len(path)-1])
		},
		GetRootNode: func() *jsonc.Node {
			root, _ := jsonc.ParseTree(document.GetText(), nil)
			return root
		},
	}
	for _, entry := range j.Entries {
		for _, jsonPath := range entry.Path {
			if jsonPath.IsKey == location.IsAtPropertyKey && location.Path.Matches(jsonPath.Path) {
				if entry.Matcher == nil || entry.Matcher(&params) {
					return &entry, &params
				}
			}
		}
	}
	return nil, nil
}

func (j *JsonHandler) isMolangLocation(location *jsonc.Location) bool {
	if location.IsAtPropertyKey || location.PreviousNode == nil {
		return false
	}
	nodeValue, ok := location.PreviousNode.Value.(string)
	if !ok {
		return false
	}
	if len(nodeValue) > 0 {
		if nodeValue[0] == '@' || nodeValue[0] == '/' {
			return false
		}
	}
	if j.MolangLocations != nil {
		for _, jsonPath := range j.MolangLocations {
			if location.Path.Matches(jsonPath.Path) {
				return true
			}
		}
	}
	return false
}

func (j *JsonHandler) isMolangSemanticLocation(location *jsonc.Location) bool {
	if location.IsAtPropertyKey || location.PreviousNode == nil {
		return false
	}
	if j.MolangSemanticLocations != nil {
		for _, jsonPath := range j.MolangSemanticLocations {
			if location.Path.Matches(jsonPath.Path) {
				return true
			}
		}
	}
	return j.isMolangLocation(location)
}

func (j *JsonHandler) commandEntry(location *jsonc.Location) *JsonCommandEntry {
	if location.IsAtPropertyKey || location.PreviousNode == nil {
		return nil
	}
	nodeValue, ok := location.PreviousNode.Value.(string)
	if !ok {
		return nil
	}
	for _, jsonPath := range j.CommandEntry.Path {
		if location.Path.Matches(jsonPath.Path) {
			if j.CommandEntry.RequireSlash {
				if len(nodeValue) == 0 || nodeValue[0] != '/' {
					return nil
				}
			}
			return &j.CommandEntry
		}
	}
	return nil
}

func (j *JsonHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode

	res := []protocol.CompletionItem{}
	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return entry.Handler.Completions(doc, position)
	}
	if j.isMolangLocation(location) {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		if isInside(docRange, position) {
			doc := document.CreateVirtualDocument(docRange)
			res = Molang.Completions(doc, position)
		}
	}
	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil {
		return res
	}

	var items []core.Symbol
	if entry.FilterDiff {
		items = difference(j.Pattern, entry.Source(ctx), entry.References(ctx))
	} else {
		items = entry.Source(ctx)
	}

	set := mapset.NewThreadUnsafeSet[string]()
	if entry.Store != nil && entry.Store.VanillaData != nil {
		set = entry.Store.VanillaData.Clone()
	}

	for _, item := range items {
		if set.ContainsOne(item.Value) {
			continue
		}
		set.Add(item.Value)
		value := `"` + item.Value + `"`
		completion := protocol.CompletionItem{
			Label: value,
		}
		if node != nil {
			completion.TextEdit = &protocol.Or_CompletionItem_textEdit{
				Value: protocol.TextEdit{
					Range: protocol.Range{
						Start: document.PositionAt(node.Offset),
						End:   document.PositionAt(node.Offset + node.Length),
					},
					NewText: value,
				},
			}
		}
		res = append(res, completion)
	}
	return res
}

func (j *JsonHandler) Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}

	res := []protocol.LocationLink{}
	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		res = entry.Handler.Definitions(doc, position)
	}
	if j.isMolangLocation(location) {
		doc := document.CreateVirtualDocument(protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		})
		res = slices.Concat(res, Molang.Definitions(doc, position))
	}
	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil {
		return res
	}

	nodeValue, ok := node.Value.(string)
	if !ok {
		return res
	}
	if entry.Transform != nil {
		nodeValue = entry.Transform(nodeValue)
	}

	for _, item := range entry.Source(ctx) {
		if item.Value != nodeValue {
			continue
		}
		location := protocol.LocationLink{
			OriginSelectionRange: &protocol.Range{
				Start: document.PositionAt(node.Offset),
				End:   document.PositionAt(node.Offset + node.Length),
			},
			TargetURI: item.URI,
		}
		if item.Range != nil {
			location.TargetRange = *item.Range
			location.TargetSelectionRange = *item.Range
		}
		res = append(res, location)
	}
	return res
}

func (j *JsonHandler) PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}

	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return entry.Handler.PrepareRename(doc, position)
	}

	entry, _ := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil || entry.DisableRename {
		return nil
	}
	// TODO: Support renaming entry that uses transform
	if entry.Transform != nil {
		return nil
	}

	start := node.Offset + 1
	end := node.Offset + node.Length - 1
	return &protocol.PrepareRenamePlaceholder{
		Range: protocol.Range{
			Start: document.PositionAt(start),
			End:   document.PositionAt(end),
		},
		Placeholder: document.GetText()[start:end],
	}
}

func (j *JsonHandler) Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if node == nil {
		return nil
	}

	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return entry.Handler.Rename(doc, position, newName)
	}

	entry, ctx := j.prepareContext(document, location)
	if entry == nil || entry.Source == nil || entry.References == nil || entry.DisableRename {
		return nil
	}

	changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
	for _, item := range slices.Concat(entry.Source(ctx), entry.References(ctx)) {
		if item.Value != node.Value {
			continue
		}
		edit := protocol.TextEdit{
			NewText: newName,
			Range:   *item.Range,
		}
		changes[item.URI] = append(changes[item.URI], edit)
	}
	return &protocol.WorkspaceEdit{Changes: changes}
}

func (j *JsonHandler) Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return entry.Handler.Hover(doc, position)
	}
	if j.isMolangLocation(location) {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return Molang.Hover(doc, position)
	}
	return nil
}

func (j *JsonHandler) SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp {
	offset := document.OffsetAt(position)
	location := jsonc.GetLocation(document.GetText(), offset)
	node := location.PreviousNode
	if entry := j.commandEntry(location); entry != nil {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return entry.Handler.SignatureHelp(doc, position)
	}
	if j.isMolangLocation(location) {
		docRange := protocol.Range{
			Start: document.PositionAt(node.Offset + 1),
			End:   document.PositionAt(node.Offset + node.Length - 1),
		}
		doc := document.CreateVirtualDocument(docRange)
		return Molang.SignatureHelp(doc, position)
	}
	return nil
}

func (j *JsonHandler) SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := []semtok.Token{}

	molangRanges := []protocol.Range{}
	commandRanges := []protocol.Range{}
	jsonc.Visit(document.GetText(), &jsonc.Visitor{
		OnLiteralValue: func(value any, offset, length, startLine, startCharacter uint32, pathSupplier func() jsonc.Path) {
			text, ok := value.(string)
			if !ok || len(text) == 0 {
				return
			}
			location := jsonc.Location{
				Path: pathSupplier(),
				PreviousNode: &jsonc.Node{
					Type:   jsonc.NodeTypeString,
					Value:  value,
					Offset: offset,
					Length: length,
				},
			}
			if entry := j.commandEntry(&location); entry != nil {
				startOffset := offset + 1
				if strings.HasPrefix(text, "/") {
					startOffset++
				} else if entry.RequireSlash {
					return
				}
				endOffset := offset + length - 1
				r := protocol.Range{
					Start: document.PositionAt(startOffset),
					End:   document.PositionAt(endOffset),
				}
				commandRanges = append(commandRanges, r)
				return
			}
			if j.isMolangSemanticLocation(&location) {
				molangRanges = append(molangRanges, protocol.Range{
					Start: document.PositionAt(offset + 1),
					End:   document.PositionAt(offset + length - 1),
				})
			}
		},
	}, nil)
	molangDocument := document.CreateVirtualDocument(molangRanges...)
	tokens = append(tokens, Molang.ComputeSemanticTokens(molangDocument)...)

	if len(commandRanges) > 0 {
		commandDocument := document.CreateVirtualDocument(commandRanges...)
		tokens = append(tokens, j.CommandEntry.Handler.ComputeSemanticTokens(commandDocument)...)
	}

	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}
