package handlers

import (
	"log"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/protocol"
	"github.com/rockide/language-server/internal/protocol/semtok"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

// TODO:
// JSON Parse
// JSON Defintions
// JSON Rename

type CommandHandler struct {
	Pattern      shared.Pattern
	Parser       *mcfunction.Parser
	EscapeQuotes bool
}

func (h *CommandHandler) GetPattern() shared.Pattern {
	return h.Pattern
}

func (h *CommandHandler) Parse(uri protocol.DocumentURI) error {
	if strings.HasSuffix(uri.Path(), ".mcfunction") {
		stores.McFunctionPath.Insert(h.Pattern, uri)
	}
	document, err := textdocument.GetOrReadFile(uri)
	if err != nil {
		return err
	}
	content := document.GetContent()
	root, _ := h.Parser.Parse(content)
	mcfunction.WalkNodeTree(root, func(i mcfunction.INode) bool {
		nodeSpec, ok := i.(mcfunction.NodeParam)
		if !ok {
			return true
		}
		paramSpec, ok := nodeSpec.ParamSpec()
		if !ok {
			return true
		}
		for _, tag := range paramSpec.Tags {
			entry, ok := commandEntries[tag]
			if !ok || entry.Store == nil {
				continue
			}
			scope := defaultScope
			if entry.Scope != nil {
				scope = entry.Scope(i)
			}
			s, e := i.Range()
			entry.Store.Insert(scope, core.Symbol{
				URI:   uri,
				Value: i.Text(content),
				Range: &protocol.Range{
					Start: document.PositionAt(s),
					End:   document.PositionAt(e),
				},
			})
		}
		return true
	})
	return nil
}

func (h *CommandHandler) Delete(uri protocol.DocumentURI) {
	if strings.HasSuffix(uri.Path(), ".mcfunction") {
		stores.McFunctionPath.Delete(uri)
	}
	for _, entry := range commandEntries {
		if entry.Store != nil {
			entry.Store.Delete(uri)
		}
	}
}

func (h *CommandHandler) Completions(document *textdocument.TextDocument, position protocol.Position) []protocol.CompletionItem {
	result := []protocol.CompletionItem{}
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	startOffset := parsed.StartOffset
	line := parsed.Content

	node := mcfunction.NodeAt(root, rOffset)
	rStart, rEnd := node.Range()
	cursorRange := protocol.Range{
		Start: position,
		End:   position,
	}
	nodeRange := protocol.Range{
		Start: document.PositionAt(startOffset + rStart),
		End:   document.PositionAt(startOffset + rEnd),
	}
	switch node.Kind() {
	case mcfunction.NodeKindFile:
		result = h.commandCompletions(cursorRange)
	case mcfunction.NodeKindCommandLit:
		result = h.commandCompletions(nodeRange)
	case mcfunction.NodeKindCommandArg:
		arg, ok := node.(mcfunction.INodeArg)
		if ok {
			switch arg.ParamKind() {
			case mcfunction.ParameterKindMap, mcfunction.ParameterKindMapPair, mcfunction.ParameterKindSelectorArg, mcfunction.ParameterKindMapJSON:
				return h.mapCompletions(arg, cursorRange, nodeRange)
			case mcfunction.ParameterKindRawMessage:
				doc := document.CreateVirtualDocument(nodeRange)
				rawMessage.CommandHandler = h
				return rawMessage.Completions(doc, position)
			case mcfunction.ParameterKindItemNbt:
				doc := document.CreateVirtualDocument(nodeRange)
				return itemNbt.Completions(doc, position)
			}
		}
		invalid := node.Parent().Kind() == mcfunction.NodeKindInvalidCommand
		if _, ok := node.(mcfunction.INodeCommand); invalid || !ok {
			root, _ := h.Parser.Parse(line[:rStart])
			n := mcfunction.NodeAt(root, rStart)
			if n, ok := n.(mcfunction.INodeCommand); ok {
				node = n
				cursorRange = nodeRange
			}
		}
		fallthrough
	case mcfunction.NodeKindCommand:
		nodeCommand, ok := node.(mcfunction.INodeCommand)
		if !ok {
			log.Printf("Failed to cast node to INodeCommand")
			return result
		}
		result = h.paramCompletions(nodeCommand, cursorRange)
	}
	return result
}

func (h *CommandHandler) commandCompletions(editRange protocol.Range) []protocol.CompletionItem {
	result := []protocol.CompletionItem{}
	set := mapset.NewThreadUnsafeSet[string]()
	for name, spec := range h.Parser.RegisteredCommands() {
		if set.ContainsOne(name) {
			continue
		}
		set.Add(name)
		result = append(result, protocol.CompletionItem{
			Label:  name,
			Detail: spec.Description,
			Kind:   protocol.MethodCompletion,
			TextEdit: &protocol.Or_CompletionItem_textEdit{
				Value: protocol.TextEdit{
					NewText: name,
					Range:   editRange,
				},
			},
		})
	}
	return result
}

func (h *CommandHandler) innerCompletions(node mcfunction.INodeCommand, paramSpec *mcfunction.ParameterSpec, editRange protocol.Range) []protocol.CompletionItem {
	result := []protocol.CompletionItem{}
	set := mapset.NewThreadUnsafeSet[string]()
	escape := func(s string) string {
		if strings.ContainsAny(s, " /") {
			s = `"` + s + `"`
		}
		if h.EscapeQuotes {
			s = strings.ReplaceAll(s, `"`, `\"`)
		}
		return s
	}
	addItem := func(value string, kind ...protocol.CompletionItemKind) {
		if set.ContainsOne(value) {
			return
		}
		set.Add(value)
		var k protocol.CompletionItemKind
		if len(kind) > 0 {
			k = kind[0]
		}
		result = append(result, protocol.CompletionItem{
			Label: value,
			Kind:  k,
			TextEdit: &protocol.Or_CompletionItem_textEdit{
				Value: protocol.TextEdit{
					NewText: value,
					Range:   editRange,
				},
			},
		})
	}
	var paramCompletions func(param *mcfunction.ParameterSpec)
	paramCompletions = func(param *mcfunction.ParameterSpec) {
		if len(param.Literals) > 0 {
			for _, lit := range param.Literals {
				addItem(escape(lit), protocol.KeywordCompletion)
			}
		}
		switch param.Kind {
		case mcfunction.ParameterKindNumber:
			addItem("0.0")
		case mcfunction.ParameterKindInteger:
			addItem("0")
		case mcfunction.ParameterKindBoolean:
			addItem("true", protocol.KeywordCompletion)
			addItem("false", protocol.KeywordCompletion)
		case mcfunction.ParameterKindSelector:
			for sel, ok := range h.Parser.GetSelectors() {
				if ok {
					addItem("@" + sel)
				}
			}
		case mcfunction.ParameterKindMap:
			result = append(result, protocol.CompletionItem{
				Label:            "[]",
				Kind:             protocol.SnippetCompletion,
				InsertTextFormat: &snippetTextFormat,
				InsertText:       "[$1=$0]",
			})
		case mcfunction.ParameterKindMapJSON, mcfunction.ParameterKindJSON:
			result = append(result, protocol.CompletionItem{
				Label:            "{}",
				Kind:             protocol.SnippetCompletion,
				InsertTextFormat: &snippetTextFormat,
				InsertText:       "{$0}",
			})
		case mcfunction.ParameterKindVector2:
			addItem("~~")
			addItem("0 0")
			addItem("^^")
		case mcfunction.ParameterKindVector3:
			addItem("~~~")
			addItem("0 0 0")
			addItem("^^^")
		case mcfunction.ParameterKindRange:
			addItem("0..0")
			addItem("..0")
			addItem("0..")
			addItem("!0..0")
		case mcfunction.ParameterKindSuffixedInteger:
			addItem("1" + param.Suffix)
		case mcfunction.ParameterKindWildcardInteger:
			addItem("*")
			addItem("0")
		case mcfunction.ParameterKindChainedCommand:
			if node != nil {
				for _, o := range node.OverloadStates() {
					paramCompletions(&o.Parameters()[0])
				}
			}
		case mcfunction.ParameterKindCommand:
			for name := range h.Parser.RegisteredCommands() {
				addItem(name, protocol.MethodCompletion)
			}
		case mcfunction.ParameterKindRawMessage:
			result = append(result, protocol.CompletionItem{
				Label:            "Raw Message Snippet...",
				Kind:             protocol.SnippetCompletion,
				InsertTextFormat: &snippetTextFormat,
				InsertText:       escape(`{"rawtext":[$0]}`),
			})
		case mcfunction.ParameterKindItemNbt:
			result = append(result, protocol.CompletionItem{
				Label:            "Item NBT Snippet...",
				Kind:             protocol.SnippetCompletion,
				InsertTextFormat: &snippetTextFormat,
				InsertText:       escape(`{$0}`),
			})
		case mcfunction.ParameterKindRelativeNumber:
			addItem("~")
			addItem("0")
		}
	}
	tagSet := mapset.NewThreadUnsafeSet[string]()
	tagCompletions := func(param *mcfunction.ParameterSpec) {
		for _, tag := range param.Tags {
			if tagSet.ContainsOne(tag) {
				continue
			}
			tagSet.Add(tag)
			entry, ok := commandEntries[tag]
			if !ok || entry.Source == nil {
				continue
			}
			for _, item := range entry.Source(node) {
				value := item.Value
				if entry.Transform != nil {
					value = entry.Transform(value)
				}
				addItem(escape(value))
			}
			if entry.Store != nil && entry.Store.VanillaData != nil {
				for value := range mapset.Elements(entry.Store.VanillaData) {
					if entry.Namespaced && !strings.HasPrefix(value, "minecraft:") {
						continue
					}
					addItem(escape(value), protocol.EnumCompletion)
				}
			}
		}
	}
	if node != nil {
		for _, o := range node.OverloadStates() {
			if !o.Matched() {
				continue
			}
			param, ok := o.Peek()
			if !ok {
				continue
			}
			paramCompletions(&param)
			tagCompletions(&param)
		}
	}
	if paramSpec != nil {
		paramCompletions(paramSpec)
		tagCompletions(paramSpec)
	}
	return result
}

func (h *CommandHandler) paramCompletions(node mcfunction.INodeCommand, editRange protocol.Range) []protocol.CompletionItem {
	return h.innerCompletions(node, nil, editRange)
}

func (h *CommandHandler) paramSpecCompletions(spec *mcfunction.ParameterSpec, editRange protocol.Range) []protocol.CompletionItem {
	return h.innerCompletions(nil, spec, editRange)
}

func (h *CommandHandler) mapCompletions(node mcfunction.INodeArg, cursorRange protocol.Range, nodeRange protocol.Range) []protocol.CompletionItem {
	keyCompletions := func(spec *mcfunction.ParameterSpec, keys []string, editRange protocol.Range) []protocol.CompletionItem {
		var result []protocol.CompletionItem
		for _, key := range keys {
			result = append(result, protocol.CompletionItem{
				Label: key,
				Kind:  protocol.FieldCompletion,
				TextEdit: &protocol.Or_CompletionItem_textEdit{
					Value: protocol.TextEdit{
						NewText: key,
						Range:   editRange,
					},
				},
			})
		}
		if spec != nil && spec.Kind != mcfunction.ParameterKindUnknown {
			result = slices.Concat(result, h.paramSpecCompletions(spec, editRange))
		}
		return result
	}
	switch n := node.(type) {
	case mcfunction.INodeArgMap:
		spec, _ := n.MapSpec().KeySpec()
		return keyCompletions(spec, n.MapSpec().Keys(), cursorRange)
	case mcfunction.INodeArgPairChild:
		kind := n.PairKind()
		switch kind {
		case mcfunction.PairKindKey:
			spec, _ := n.KeySpec()
			keys := n.Keys()
			return keyCompletions(&spec, keys, nodeRange)
		case mcfunction.PairKindEqual, mcfunction.PairKindValue:
			spec, ok := n.ValueSpec()
			if ok {
				if kind == mcfunction.PairKindValue {
					cursorRange = nodeRange
				}
				return h.paramSpecCompletions(&spec, cursorRange)
			}
		}
	}
	return []protocol.CompletionItem{}
}

func (h *CommandHandler) Definitions(document *textdocument.TextDocument, position protocol.Position) []protocol.LocationLink {
	result := []protocol.LocationLink{}
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	startOffset := parsed.StartOffset
	line := parsed.Content
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	nodeSpec, ok := node.(mcfunction.NodeParam)
	if !ok {
		return result
	}
	paramSpec, ok := nodeSpec.ParamSpec()
	if !ok {
		return result
	}
	rStart, rEnd := node.Range()
	nodeRange := protocol.Range{
		Start: document.PositionAt(startOffset + rStart),
		End:   document.PositionAt(startOffset + rEnd),
	}
	nodeValue := node.Text(line)
	if h.EscapeQuotes {
		nodeValue = strings.Trim(nodeValue, `\"`)
	} else {
		nodeValue = strings.Trim(nodeValue, `"`)
	}
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok || entry.Source == nil {
			continue
		}
		for _, item := range entry.Source(node) {
			value := item.Value
			if entry.Transform != nil {
				value = entry.Transform(value)
			}
			if value != nodeValue {
				continue
			}
			location := protocol.LocationLink{
				OriginSelectionRange: &nodeRange,
				TargetURI:            item.URI,
			}
			if item.Range != nil {
				location.TargetRange = *item.Range
				location.TargetSelectionRange = *item.Range
			}
			result = append(result, location)
		}
	}
	return result
}

func (h *CommandHandler) PrepareRename(document *textdocument.TextDocument, position protocol.Position) *protocol.PrepareRenamePlaceholder {
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	startOffset := parsed.StartOffset
	line := parsed.Content
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	nodeSpec, ok := node.(mcfunction.NodeParam)
	if !ok {
		return nil
	}
	paramSpec, ok := nodeSpec.ParamSpec()
	if !ok {
		return nil
	}
	rStart, rEnd := node.Range()
	nodeRange := protocol.Range{
		Start: document.PositionAt(startOffset + rStart),
		End:   document.PositionAt(startOffset + rEnd),
	}
	nodeValue := node.Text(line)
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok || entry.DisableRename {
			continue
		}
		return &protocol.PrepareRenamePlaceholder{
			Range:       nodeRange,
			Placeholder: nodeValue,
		}
	}
	return nil
}

func (h *CommandHandler) Rename(document *textdocument.TextDocument, position protocol.Position, newName string) *protocol.WorkspaceEdit {
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	line := parsed.Content
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	nodeSpec, ok := node.(mcfunction.NodeParam)
	if !ok {
		return nil
	}
	paramSpec, ok := nodeSpec.ParamSpec()
	if !ok {
		return nil
	}
	nodeValue := node.Text(line)
	// TODO: Check cross rename between unescaped and escaped strings
	changes := make(map[protocol.DocumentURI][]protocol.TextEdit)
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok || entry.DisableRename {
			continue
		}
		items := entry.Source(node)
		if entry.References != nil {
			items = slices.Concat(items, entry.References(node))
		}
		for _, item := range items {
			if item.Value != nodeValue {
				continue
			}
			edit := protocol.TextEdit{
				NewText: newName,
				Range:   *item.Range,
			}
			changes[item.URI] = append(changes[item.URI], edit)
		}
	}
	return &protocol.WorkspaceEdit{Changes: changes}
}

func (h *CommandHandler) Hover(document *textdocument.TextDocument, position protocol.Position) *protocol.Hover {
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	var commandNode mcfunction.INodeCommand
	switch n := node.(type) {
	case mcfunction.INodeCommand:
		commandNode = n
	case mcfunction.INodeArg:
		if parent, ok := n.Parent().(mcfunction.INodeCommand); ok {
			commandNode = parent
		}
	}
	if commandNode == nil {
		return nil
	}
	spec := commandNode.Spec()
	if spec == nil {
		return nil
	}
	return &protocol.Hover{
		Contents: protocol.MarkupContent{
			Kind: protocol.Markdown,
			Value: "```\n" +
				commandNode.CommandName() +
				"\n```\n" +
				spec.Description,
		},
	}
}

func (h *CommandHandler) SignatureHelp(document *textdocument.TextDocument, position protocol.Position) *protocol.SignatureHelp {
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	line := parsed.Content
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	var commandNode mcfunction.INodeCommand
	flag := false
REPARSE:
	rStart, _ := node.Range()
	switch n := node.(type) {
	case mcfunction.INodeCommand:
		commandNode = n
	case mcfunction.INodeArg:
		if parent, ok := n.Parent().(mcfunction.INodeCommand); ok {
			if parent.IsValid() {
				commandNode = parent
			} else if !flag {
				root, _ := h.Parser.Parse(line[:rStart])
				node = mcfunction.NodeAt(root, rStart)
				flag = true
				goto REPARSE
			}
		}
	}
	if commandNode == nil {
		return nil
	}
	spec := commandNode.Spec()
	signatures := make([]protocol.SignatureInformation, 0, len(commandNode.OverloadStates()))
	for _, o := range commandNode.OverloadStates() {
		if !o.Matched() {
			continue
		}
		params := o.Parameters()
		index := uint32(o.Index())
		paramLabels := make([]string, len(params))
		for j, p := range params {
			paramLabels[j] = p.ToString()
		}
		signatures = append(signatures, protocol.SignatureInformation{
			Label: commandNode.CommandName() + " " + strings.Join(paramLabels, " "),
			Documentation: &protocol.Or_SignatureInformation_documentation{
				Value: spec.Description,
			},
			Parameters: sliceutil.Map(paramLabels, func(label string) protocol.ParameterInformation {
				return protocol.ParameterInformation{
					Label: label,
				}
			}),
			ActiveParameter: index,
		})
	}
	return &protocol.SignatureHelp{
		Signatures: signatures,
	}
}

func (h *CommandHandler) ComputeSemanticTokens(document *textdocument.TextDocument) []semtok.Token {
	content := document.GetContent()
	root, _ := h.Parser.Parse(content)
	tokens := []semtok.Token{}
	molangRanges := []protocol.Range{}
	isMolang := func(node mcfunction.INodeArg) bool {
		param, ok := node.ParamSpec()
		if ok && slices.Contains(param.Tags, mcfunction.TagMolang) {
			start, end := node.Range()
			length := end - start
			if length >= 2 {
				if content[start] == '"' && content[end-1] == '"' {
					return true
				}
			}
		}
		return false
	}
	mcfunction.WalkNodeTree(root, func(i mcfunction.INode) bool {
		start, end := i.Range()
		length := end - start
		pA := document.PositionAt(start)
		pB := document.PositionAt(end)
		switch n := i.(type) {
		case mcfunction.INodeCommand:
			if p, ok := i.(mcfunction.INodeArg); n.Kind() != mcfunction.NodeKindCommandArg || (ok && p.ParamKind() == mcfunction.ParameterKindCommand) {
				tokens = append(tokens, semtok.Token{
					Type:  semtok.TokNamespace,
					Line:  pA.Line,
					Start: pA.Character,
					Len:   uint32(len(n.CommandName())),
				})
			}
		case mcfunction.INodeArg:
			if isMolang(n) {
				tokens = append(tokens, semtok.Token{
					Type:  semtok.TokString,
					Line:  pA.Line,
					Start: pA.Character,
					Len:   1,
				})
				tokens = append(tokens, semtok.Token{
					Type:  semtok.TokString,
					Line:  pB.Line,
					Start: pB.Character - 1,
					Len:   1,
				})
				pA.Character += 1
				pB.Character -= 1
				molangRanges = append(molangRanges, protocol.Range{
					Start: pA,
					End:   pB,
				})
			} else if tokType, ok := commandParamTokenMap[n.ParamKind()]; ok {
				spec, ok := n.ParamSpec()
				if ok && slices.Contains(spec.Tags, mcfunction.TagExecuteChain) {
					tokType = semtok.TokKeyword
				}
				tokens = append(tokens, semtok.Token{
					Type:  tokType,
					Line:  pA.Line,
					Start: pA.Character,
					Len:   length,
				})
			}
		}
		return true
	})
	molangDocument := document.CreateVirtualDocument(molangRanges...)
	tokens = slices.Concat(tokens, Molang.ComputeSemanticTokens(molangDocument))
	return tokens
}

func (h *CommandHandler) SemanticTokens(document *textdocument.TextDocument) *protocol.SemanticTokens {
	tokens := h.ComputeSemanticTokens(document)
	return &protocol.SemanticTokens{
		Data: semtok.Encode(tokens, tokenType, tokenModifier),
	}
}

func (h *CommandHandler) parseLine(document *textdocument.TextDocument, position protocol.Position) parsedCommandLine {
	content := document.GetContent()
	startOffset := document.OffsetAt(protocol.Position{
		Line: position.Line,
	})
	offset := document.OffsetAt(position)
	rOffset := offset - startOffset
	line := content[startOffset:]
	root, _ := h.Parser.Parse(line)
	return parsedCommandLine{
		Content:        line,
		StartOffset:    startOffset,
		RelativeOffset: rOffset,
		Root:           root,
	}
}

type parsedCommandLine struct {
	Content        []rune
	StartOffset    uint32
	RelativeOffset uint32
	Root           mcfunction.INode
}

var commandParamTokenMap = map[mcfunction.ParameterKind]semtok.Type{
	mcfunction.ParameterKindLiteral:         semtok.TokMethod,
	mcfunction.ParameterKindString:          semtok.TokString,
	mcfunction.ParameterKindNumber:          semtok.TokNumber,
	mcfunction.ParameterKindInteger:         semtok.TokNumber,
	mcfunction.ParameterKindBoolean:         semtok.TokMacro,
	mcfunction.ParameterKindSelector:        semtok.TokEnumMember,
	mcfunction.ParameterKindMap:             semtok.TokEnumMember,
	mcfunction.ParameterKindRange:           semtok.TokNumber,
	mcfunction.ParameterKindSuffixedInteger: semtok.TokNumber,
	mcfunction.ParameterKindWildcardInteger: semtok.TokNumber,
	mcfunction.ParameterKindChainedCommand:  semtok.TokKeyword,
	mcfunction.ParameterKindRawMessage:      semtok.TokString,
	mcfunction.ParameterKindItemNbt:         semtok.TokString,
	mcfunction.ParameterKindRelativeNumber:  semtok.TokNumber,
}

type commandEntry struct {
	Store         *stores.SymbolStore
	Source        func(node mcfunction.INode) []core.Symbol
	References    func(node mcfunction.INode) []core.Symbol
	Transform     func(s string) string
	Scope         func(node mcfunction.INode) string
	Namespaced    bool
	DisableRename bool
}

var commandEntries = map[string]commandEntry{
	mcfunction.TagAimAssistId: {
		Store: stores.AimAssistCategory.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.AimAssistCategory.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.AimAssistCategory.References.Get()
		},
	},
	mcfunction.TagBiomeId: {
		Store: stores.BiomeId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.BiomeId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.BiomeId.References.Get()
		},
	},
	mcfunction.TagBlockId: {
		Store: stores.ItemId.References,
		Scope: func(node mcfunction.INode) string {
			return "block"
		},
		Namespaced: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.ItemId.Source.Get("block")
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.ItemId.References.Get("block")
		},
	},
	mcfunction.TagBlockState: {
		Store: stores.BlockState.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.BlockState.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.BlockState.References.Get()
		},
	},
	mcfunction.TagCameraId: {
		Store: stores.CameraId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.CameraId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.CameraId.References.Get()
		},
	},
	mcfunction.TagClientAnimationId: {
		Store: stores.ClientAnimation.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.ClientAnimation.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.ClientAnimation.References.Get()
		},
	},
	mcfunction.TagDialogueId: {
		Store: stores.DialogueId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.DialogueId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.DialogueId.References.Get()
		},
	},
	mcfunction.TagEntityEvent: {
		DisableRename: true,
		Store:         stores.EntityEvent.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.EntityEvent.Source.Get()
		},
	},
	mcfunction.TagEntityId: {
		Store:      stores.EntityId.References,
		Namespaced: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.EntityId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.EntityId.References.Get()
		},
	},
	mcfunction.TagFeatureId: {
		Store: stores.FeatureId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.FeatureId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.FeatureId.References.Get()
		},
	},
	mcfunction.TagFeatureRuleId: {
		Store: stores.FeatureRuleId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.FeatureRuleId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.FeatureRuleId.References.Get()
		},
	},
	mcfunction.TagFogId: {
		Store: stores.Fog.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.Fog.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.Fog.References.Get()
		},
	},
	mcfunction.TagFunctionFile: {
		DisableRename: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.McFunctionPath.Get()
		},
		Transform: func(s string) string {
			return strings.TrimPrefix(s, "functions/")
		},
	},
	mcfunction.TagItemId: {
		Store:      stores.ItemId.References,
		Namespaced: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.ItemId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.ItemId.References.Get()
		},
	},
	mcfunction.TagJigsawId: {
		Store: stores.WorldgenJigsaw.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.WorldgenJigsaw.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.WorldgenJigsaw.References.Get()
		},
	},
	mcfunction.TagJigsawTemplatePoolId: {
		Store: stores.WorldgenTemplatePool.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.WorldgenTemplatePool.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.WorldgenTemplatePool.References.Get()
		},
	},
	mcfunction.TagLootTableFile: {
		DisableRename: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.LootTablePath.Get()
		},
		Transform: func(s string) string {
			s = strings.TrimPrefix(s, "loot_tables/")
			s = strings.TrimSuffix(s, ".json")
			return s
		},
	},
	mcfunction.TagMusicId: {
		Store: stores.MusicDefinition.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.MusicDefinition.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.MusicDefinition.References.Get()
		},
	},
	mcfunction.TagProvidedFogId: {
		Store: stores.ProvidedFogId.Source,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.ProvidedFogId.Source.Get()
		},
	},
	mcfunction.TagRecipeId: {
		Store: stores.RecipeId.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.RecipeId.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.RecipeId.References.Get()
		},
	},
	mcfunction.TagSoundId: {
		Store: stores.SoundDefinition.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.SoundDefinition.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.SoundDefinition.References.Get()
		},
	},
	mcfunction.TagStructureFile: {
		DisableRename: true,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.StructurePath.Get()
		},
	},
	mcfunction.TagTagId: {
		Store: stores.Tag.Source,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.Tag.Source.Get()
		},
	},
	mcfunction.TagTickingAreaId: {
		Store: stores.TickingArea.Source,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.TickingArea.Source.Get()
		},
	},
	mcfunction.TagTypeFamilyId: {
		Store: stores.EntityFamily.References,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.EntityFamily.Source.Get()
		},
		References: func(node mcfunction.INode) []core.Symbol {
			return stores.EntityFamily.References.Get()
		},
	},
	mcfunction.TagScoreboardObjectiveId: {
		Store: stores.ScoreboardObjective.Source,
		Source: func(node mcfunction.INode) []core.Symbol {
			return stores.ScoreboardObjective.Source.Get()
		},
	},
}

var rawMessage = &JsonRawMessageHandler{
	JsonSnippetHandler: &JsonSnippetHandler{
		JsonHandler: &JsonHandler{
			Entries: []JsonEntry{
				{
					Store: stores.Lang.References,
					Path:  []shared.JsonPath{shared.JsonValue("rawtext/**/translate")},
					Source: func(ctx *JsonContext) []core.Symbol {
						return stores.Lang.Source.Get()
					},
					References: func(ctx *JsonContext) []core.Symbol {
						return stores.Lang.References.Get()
					},
				},
				{
					Store: stores.ScoreboardObjective.Source,
					Path:  []shared.JsonPath{shared.JsonValue("rawtext/**/score/objective")},
					Source: func(ctx *JsonContext) []core.Symbol {
						return stores.ScoreboardObjective.Source.Get()
					},
					References: func(ctx *JsonContext) []core.Symbol {
						return nil
					},
				},
			},
		},
		SnippetEntries: []SnippetEntry{
			{
				Path: []shared.JsonPath{shared.JsonValue("**/rawtext/*")},
				Snippets: []Snippet{
					{
						Label: "Text",
						Value: `{"text":"$0"}`,
					},
					{
						Label: "Score",
						Value: `{"score":{"name":"$0","objective":"$1"}}`,
					},
					{
						Label: "Translate",
						Value: `{"translate":"$0"}`,
					},
					{
						Label: "Translate with",
						Value: `{"translate":"$1","with":{"rawtext":[$0]}}`,
					},
				},
			},
			{
				Path: []shared.JsonPath{shared.JsonKey("**/rawtext/**")},
				Snippets: []Snippet{
					{
						Label: "With",
						Value: `"with":{"rawtext":[$0]}`,
					},
				},
			},
			{
				Path: []shared.JsonPath{shared.JsonValue("**/rawtext/**/with/*")},
				Snippets: []Snippet{
					{
						Label: "Rawtext",
						Value: `{"rawtext":[$0]}`,
					},
				},
			},
		},
	},
}

var itemNbt = &JsonSnippetHandler{
	JsonHandler: &JsonHandler{
		Entries: []JsonEntry{
			{
				Store: stores.ItemId.References,
				Path: []shared.JsonPath{
					shared.JsonValue("minecraft:can_place_on/blocks/*"),
					shared.JsonValue("minecraft:can_destroy/blocks/*"),
				},
				ScopeKey: func(ctx *JsonContext) string {
					return "block"
				},
				Source: func(ctx *JsonContext) []core.Symbol {
					return stores.ItemId.Source.Get("block")
				},
				References: func(ctx *JsonContext) []core.Symbol {
					return stores.ItemId.References.Get("block")
				},
			},
		},
	},
	SnippetEntries: []SnippetEntry{
		{
			Path: []shared.JsonPath{shared.JsonKey("*")},
			Snippets: []Snippet{
				{
					Label: "Can place on",
					Value: `"minecraft:can_place_on":{"blocks":[$0]}`,
				},
				{
					Label: "Can destroy",
					Value: `"minecraft:can_destroy":{"blocks":[$0]}`,
				},
				{
					Label: "Keep on death",
					Value: `"minecraft:keep_on_death":{}`,
				},
				{
					Label: "Lock in inventory",
					Value: `"minecraft:item_lock":{"mode":"lock_in_inventory"}"`,
				},
				{
					Label: "Lock in slot",
					Value: `"minecraft:item_lock":{"mode":"lock_in_slot"}"`,
				},
			},
		},
	},
}
