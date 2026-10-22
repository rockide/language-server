package handlers

import (
	"log"
	"slices"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/mcfunction"
	"github.com/rockide/language-server/internal/semtok"
	"github.com/rockide/language-server/internal/sliceutil"
	"github.com/rockide/language-server/internal/textdocument"
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type CommandHandler struct {
	Pattern shared.Pattern
	Parser  *mcfunction.Parser
}

func (h *CommandHandler) GetPattern() shared.Pattern {
	return h.Pattern
}

func (h *CommandHandler) Parse(uri uri.URI) error {
	if strings.HasSuffix(uri.FsPath(), ".mcfunction") {
		stores.McFunctionPath.Insert(h.Pattern, uri)
	}
	document, err := textdocument.GetOrReadFile(uri)
	if err != nil {
		return err
	}
	return h.ParseDocument(document)
}

func (h *CommandHandler) ParseDocument(document *textdocument.TextDocument) error {
	content := document.GetContent()
	root, _ := h.Parser.Parse(content)
	mcfunction.WalkNodeTree(root, func(i mcfunction.Node) bool {
		nodeSpec, ok := i.(mcfunction.ParamNode)
		if !ok {
			return true
		}
		paramSpec, ok := nodeSpec.ParamSpec()
		if !ok {
			return true
		}
		for _, tag := range paramSpec.Tags {
			entry, ok := commandEntries[tag]
			if !ok {
				continue
			}
			store := entry.store()
			if store == nil {
				continue
			}
			scope := defaultScope
			if entry.Scope != nil {
				scope = entry.Scope(i)
			}
			s, e := i.Range()
			nodeValue := i.Text(content)
			if h.Parser.EscapeQuotes() {
				nodeValue = strings.ReplaceAll(nodeValue, `\"`, `"`)
			}
			store.Insert(scope, core.Symbol{
				URI:   document.URI,
				Value: nodeValue,
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

func (h *CommandHandler) Delete(uri uri.URI) {
	if strings.HasSuffix(uri.FsPath(), ".mcfunction") {
		stores.McFunctionPath.Delete(uri)
	}
	for _, entry := range commandEntries {
		if store := entry.store(); store != nil {
			store.Delete(uri)
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
		result = h.commandCompletions(cursorRange, "")
	case mcfunction.NodeKindCommandLit:
		prefix := ""
		if node.Text(line)[0] == '/' {
			prefix = "/"
		}
		result = h.commandCompletions(nodeRange, prefix)
	case mcfunction.NodeKindCommandArg:
		arg, ok := node.(mcfunction.ArgNode)
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
		if _, ok := node.(mcfunction.CommandNode); invalid || !ok {
			root, _ := h.Parser.Parse(line[:rStart])
			n := mcfunction.NodeAt(root, rStart)
			if n, ok := n.(mcfunction.CommandNode); ok {
				node = n
				cursorRange = nodeRange
			}
		}
		fallthrough
	case mcfunction.NodeKindCommand:
		nodeCommand, ok := node.(mcfunction.CommandNode)
		if !ok {
			log.Printf("Failed to cast node to CommandNode")
			return result
		}
		result = h.paramCompletions(nodeCommand, cursorRange)
	}
	return result
}

func (h *CommandHandler) commandCompletions(editRange protocol.Range, prefix string) []protocol.CompletionItem {
	result := []protocol.CompletionItem{}
	set := mapset.NewThreadUnsafeSet[string]()
	for name, spec := range h.Parser.RegisteredCommands() {
		if set.ContainsOne(name) {
			continue
		}
		set.Add(name)
		value := prefix + name
		result = append(result, protocol.CompletionItem{
			Label:  value,
			Detail: protocol.NewOptional(spec.Description),
			Kind:   protocol.CompletionItemKindMethod,
			TextEdit: &protocol.TextEdit{
				NewText: value,
				Range:   editRange,
			},
		})
	}
	return result
}

func (h *CommandHandler) innerCompletions(node mcfunction.CommandNode, paramSpec *mcfunction.ParameterSpec, editRange protocol.Range) []protocol.CompletionItem {
	result := []protocol.CompletionItem{}
	set := mapset.NewThreadUnsafeSet[string]()
	escape := func(s string) string {
		if strings.ContainsAny(s, " /") {
			s = `"` + s + `"`
		}
		if h.Parser.EscapeQuotes() {
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
		label := value
		if strings.HasPrefix(value, `\"`) {
			label = strings.ReplaceAll(value, `\"`, `"`)
		}
		result = append(result, protocol.CompletionItem{
			Label: label,
			Kind:  k,
			TextEdit: &protocol.TextEdit{
				NewText: value,
				Range:   editRange,
			},
		})
	}
	var paramCompletions func(param *mcfunction.ParameterSpec)
	paramCompletions = func(param *mcfunction.ParameterSpec) {
		if len(param.Literals) > 0 {
			for _, lit := range param.Literals {
				addItem(escape(lit), protocol.CompletionItemKindKeyword)
			}
		}
		switch param.Kind {
		case mcfunction.ParameterKindNumber:
			addItem("0.0")
		case mcfunction.ParameterKindInteger:
			addItem("0")
		case mcfunction.ParameterKindBoolean:
			addItem("true", protocol.CompletionItemKindKeyword)
			addItem("false", protocol.CompletionItemKindKeyword)
		case mcfunction.ParameterKindSelector:
			for sel, ok := range h.Parser.GetSelectors() {
				if ok {
					addItem("@" + sel)
				}
			}
		case mcfunction.ParameterKindMap:
			result = append(result, protocol.CompletionItem{
				Label:            "[]",
				Kind:             protocol.CompletionItemKindSnippet,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
				InsertText:       protocol.NewOptional("[$1=$0]"),
			})
		case mcfunction.ParameterKindMapJSON, mcfunction.ParameterKindJSON:
			result = append(result, protocol.CompletionItem{
				Label:            "{}",
				Kind:             protocol.CompletionItemKindSnippet,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
				InsertText:       protocol.NewOptional("{$0}"),
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
				addItem(name, protocol.CompletionItemKindMethod)
			}
		case mcfunction.ParameterKindRawMessage:
			result = append(result, protocol.CompletionItem{
				Label:            "Raw Message Snippet...",
				Kind:             protocol.CompletionItemKindSnippet,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
				InsertText:       protocol.NewOptional(escape(`{"rawtext":[$0]}`)),
			})
		case mcfunction.ParameterKindItemNbt:
			result = append(result, protocol.CompletionItem{
				Label:            "Item NBT Snippet...",
				Kind:             protocol.CompletionItemKindSnippet,
				InsertTextFormat: protocol.InsertTextFormatSnippet,
				InsertText:       protocol.NewOptional(escape(`{$0}`)),
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
			if !ok {
				continue
			}
			candidates := entry.Definition
			if entry.Kind == Definition {
				candidates = entry.References
			}
			for _, item := range candidates(node) {
				value := item.Value
				if entry.Transform != nil {
					value = entry.Transform(value)
				}
				addItem(escape(value))
			}
			if entry.Binding != nil && entry.Binding.VanillaData != nil {
				for value := range mapset.Elements(entry.Binding.VanillaData) {
					if entry.Namespaced && !strings.HasPrefix(value, "minecraft:") {
						continue
					}
					addItem(escape(value), protocol.CompletionItemKindEnum)
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

func (h *CommandHandler) paramCompletions(node mcfunction.CommandNode, editRange protocol.Range) []protocol.CompletionItem {
	return h.innerCompletions(node, nil, editRange)
}

func (h *CommandHandler) paramSpecCompletions(spec *mcfunction.ParameterSpec, editRange protocol.Range) []protocol.CompletionItem {
	return h.innerCompletions(nil, spec, editRange)
}

func (h *CommandHandler) mapCompletions(node mcfunction.ArgNode, cursorRange protocol.Range, nodeRange protocol.Range) []protocol.CompletionItem {
	keyCompletions := func(spec *mcfunction.ParameterSpec, keys []string, editRange protocol.Range) []protocol.CompletionItem {
		var result []protocol.CompletionItem
		for _, key := range keys {
			result = append(result, protocol.CompletionItem{
				Label: key,
				Kind:  protocol.CompletionItemKindField,
				TextEdit: &protocol.TextEdit{
					NewText: key,
					Range:   editRange,
				},
			})
		}
		if spec != nil && spec.Kind != mcfunction.ParameterKindUnknown {
			result = slices.Concat(result, h.paramSpecCompletions(spec, editRange))
		}
		return result
	}
	switch n := node.(type) {
	case mcfunction.MapNode:
		spec, _ := n.MapSpec().KeySpec()
		return keyCompletions(spec, n.MapSpec().Keys(), cursorRange)
	case mcfunction.PairChildNode:
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
	nodeSpec, ok := node.(mcfunction.ParamNode)
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
	if h.Parser.EscapeQuotes() {
		nodeValue = strings.Trim(nodeValue, `\"`)
	} else {
		nodeValue = strings.Trim(nodeValue, `"`)
	}
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok {
			continue
		}
		for _, item := range entry.Definition(node) {
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

func (h *CommandHandler) References(document *textdocument.TextDocument, position protocol.Position, includeDeclaration bool) []protocol.Location {
	result := []protocol.Location{}
	parsed := h.parseLine(document, position)
	root := parsed.Root
	rOffset := parsed.RelativeOffset
	line := parsed.Content
	if root == nil {
		return nil
	}
	node := mcfunction.NodeAt(root, rOffset)
	nodeSpec, ok := node.(mcfunction.ParamNode)
	if !ok {
		return result
	}
	paramSpec, ok := nodeSpec.ParamSpec()
	if !ok {
		return result
	}
	nodeValue := node.Text(line)
	if h.Parser.EscapeQuotes() {
		nodeValue = strings.Trim(nodeValue, `\"`)
	} else {
		nodeValue = strings.Trim(nodeValue, `"`)
	}
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok {
			continue
		}
		symbols := entry.References(node)
		if includeDeclaration {
			symbols = slices.Concat(entry.Definition(node), symbols)
		}
		for _, item := range symbols {
			value := item.Value
			if entry.Transform != nil {
				value = entry.Transform(value)
			}
			if value != nodeValue {
				continue
			}
			location := protocol.Location{
				URI: item.URI,
			}
			if item.Range != nil {
				location.Range = *item.Range
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
	nodeSpec, ok := node.(mcfunction.ParamNode)
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
	if h.Parser.EscapeQuotes() {
		nodeValue = strings.ReplaceAll(nodeValue, `\"`, `"`)
	}
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
	nodeSpec, ok := node.(mcfunction.ParamNode)
	if !ok {
		return nil
	}
	paramSpec, ok := nodeSpec.ParamSpec()
	if !ok {
		return nil
	}
	nodeValue := node.Text(line)
	if h.Parser.EscapeQuotes() {
		nodeValue = strings.ReplaceAll(nodeValue, `\"`, `"`)
	}
	changes := make(map[uri.URI][]protocol.TextEdit)
	for _, tag := range paramSpec.Tags {
		entry, ok := commandEntries[tag]
		if !ok || entry.DisableRename {
			continue
		}
		items := slices.Concat(entry.Definition(node), entry.References(node))
		for _, item := range items {
			if item.Value != nodeValue {
				continue
			}
			var newText string
			if strings.HasSuffix(item.URI.FsPath(), ".json") {
				newText = strings.ReplaceAll(newName, `"`, `\"`)
			} else {
				newText = newName
			}
			edit := protocol.TextEdit{
				NewText: newText,
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
	var commandNode mcfunction.CommandNode
	switch n := node.(type) {
	case mcfunction.CommandNode:
		commandNode = n
	case mcfunction.ArgNode:
		if parent := n.CommandNode(); parent != nil && parent.IsValid() {
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
		Contents: &protocol.MarkupContent{
			Kind: protocol.MarkupKindMarkdown,
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
	var commandNode mcfunction.CommandNode
	switch n := node.(type) {
	case mcfunction.CommandNode:
		commandNode = n
	case mcfunction.ArgNode:
		if parent := n.CommandNode(); parent != nil && parent.IsValid() {
			commandNode = parent
		} else {
			s, _ := n.Range()
			root, _ := h.Parser.Parse(line[:s])
			node = mcfunction.NodeAt(root, s)
			if n, ok := node.(mcfunction.CommandNode); ok {
				commandNode = n
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
			Label:         commandNode.CommandName() + " " + strings.Join(paramLabels, " "),
			Documentation: protocol.String(spec.Description),
			Parameters: sliceutil.Map(paramLabels, func(label string) protocol.ParameterInformation {
				return protocol.ParameterInformation{
					Label: protocol.String(label),
				}
			}),
			ActiveParameter: protocol.NewNullable(index),
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
	isMolang := func(node mcfunction.ArgNode) bool {
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
	mcfunction.WalkNodeTree(root, func(i mcfunction.Node) bool {
		start, end := i.Range()
		length := end - start
		pA := document.PositionAt(start)
		pB := document.PositionAt(end)
		switch n := i.(type) {
		case mcfunction.CommandNode:
			if p, ok := i.(mcfunction.ArgNode); n.Kind() != mcfunction.NodeKindCommandArg || (ok && p.ParamKind() == mcfunction.ParameterKindCommand) {
				if start < uint32(len(content)) && content[start] == '/' {
					tokens = append(tokens, semtok.Token{
						Type:  semtok.TokOperator,
						Line:  pA.Line,
						Start: pA.Character,
						Len:   1,
					})
					pA.Character += 1
				}
				tokens = append(tokens, semtok.Token{
					Type:  semtok.TokNamespace,
					Line:  pA.Line,
					Start: pA.Character,
					Len:   uint32(len(n.CommandName())),
				})
			}
		case mcfunction.ArgNode:
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
	endOffset := document.OffsetAt(protocol.Position{
		Line: position.Line + 1,
	})
	offset := document.OffsetAt(position)
	rOffset := offset - startOffset
	line := content[startOffset:endOffset]
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
	Root           mcfunction.Node
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
	Kind    BindingKind
	Binding *stores.SymbolBinding
	// Overrides the default Definition method
	DefinitionOverride func(node mcfunction.Node) []core.Symbol
	// Overrides the default References method
	ReferencesOverride func(node mcfunction.Node) []core.Symbol
	Transform          func(s string) string
	Scope              func(node mcfunction.Node) string
	Namespaced         bool
	DisableRename      bool
}

func (e *commandEntry) store() *stores.SymbolStore {
	if e.Binding == nil {
		return nil
	}
	switch e.Kind {
	case Definition:
		return e.Binding.Definition
	case References:
		return e.Binding.References
	default:
		return nil
	}
}

// Definition returns the declarations. Defaults to Binding.Definition.
func (e *commandEntry) Definition(node mcfunction.Node) []core.Symbol {
	if e.DefinitionOverride != nil {
		return e.DefinitionOverride(node)
	}
	if e.Binding == nil {
		return nil
	}
	return e.Binding.Definition.Get()
}

// References returns the usages. Defaults to Binding.References.
func (e *commandEntry) References(node mcfunction.Node) []core.Symbol {
	if e.ReferencesOverride != nil {
		return e.ReferencesOverride(node)
	}
	if e.Binding == nil {
		return nil
	}
	return e.Binding.References.Get()
}

var commandEntries = map[string]commandEntry{
	mcfunction.TagAimAssistId: {
		Kind:    References,
		Binding: stores.AimAssistCategory,
	},
	mcfunction.TagBiomeId: {
		Kind:    References,
		Binding: stores.BiomeId,
	},
	mcfunction.TagBlockId: {
		Kind:    References,
		Binding: stores.ItemId,
		Scope: func(node mcfunction.Node) string {
			return "block"
		},
		Namespaced: true,
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.ItemId.Definition.Get("block")
		},
		ReferencesOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.ItemId.References.Get("block")
		},
	},
	mcfunction.TagBlockState: {
		Kind:    References,
		Binding: stores.BlockState,
	},
	mcfunction.TagCameraId: {
		Kind:    References,
		Binding: stores.CameraId,
	},
	mcfunction.TagClientAnimationId: {
		Kind:    References,
		Binding: stores.ClientAnimationId,
	},
	mcfunction.TagDialogueId: {
		Kind:    References,
		Binding: stores.DialogueId,
	},
	mcfunction.TagEntityEvent: {
		DisableRename: true,
		Kind:          References,
		Binding:       stores.EntityEvent,
	},
	mcfunction.TagEntityId: {
		Kind:       References,
		Binding:    stores.EntityId,
		Namespaced: true,
	},
	mcfunction.TagFeatureId: {
		Kind:    References,
		Binding: stores.FeatureId,
	},
	mcfunction.TagFeatureRuleId: {
		Kind:    References,
		Binding: stores.FeatureRuleId,
	},
	mcfunction.TagFogId: {
		Kind:    References,
		Binding: stores.FogId,
	},
	mcfunction.TagFunctionFile: {
		DisableRename: true,
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.McFunctionPath.Get()
		},
		Transform: func(s string) string {
			return strings.TrimPrefix(s, "functions/")
		},
	},
	mcfunction.TagItemId: {
		Kind:       References,
		Binding:    stores.ItemId,
		Namespaced: true,
	},
	mcfunction.TagJigsawId: {
		Kind:    References,
		Binding: stores.WorldgenJigsawId,
	},
	mcfunction.TagJigsawTemplatePoolId: {
		Kind:    References,
		Binding: stores.WorldgenTemplatePoolId,
	},
	mcfunction.TagLootTableFile: {
		DisableRename: true,
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.LootTablePath.Get()
		},
		Transform: func(s string) string {
			s = strings.TrimPrefix(s, "loot_tables/")
			s = strings.TrimSuffix(s, ".json")
			return s
		},
	},
	mcfunction.TagMusicId: {
		Kind:    References,
		Binding: stores.MusicDefinitionId,
	},
	mcfunction.TagProvidedFogId: {
		Kind:    Definition,
		Binding: stores.ProvidedFogId,
		// ReferencesOverride already returns Binding.Definition; the default would duplicate rename edits
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return nil
		},
		ReferencesOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.ProvidedFogId.Definition.Get()
		},
	},
	mcfunction.TagRecipeId: {
		Kind:    References,
		Binding: stores.RecipeId,
	},
	mcfunction.TagSoundId: {
		Kind:    References,
		Binding: stores.SoundDefinitionId,
	},
	mcfunction.TagStructureFile: {
		DisableRename: true,
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.StructurePath.Get()
		},
	},
	mcfunction.TagTagId: {
		Kind:    Definition,
		Binding: stores.Tag,
		// ReferencesOverride already returns Binding.Definition; the default would duplicate rename edits
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return nil
		},
		ReferencesOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.Tag.Definition.Get()
		},
	},
	mcfunction.TagTickingAreaId: {
		Kind:    Definition,
		Binding: stores.TickingArea,
		// ReferencesOverride already returns Binding.Definition; the default would duplicate rename edits
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return nil
		},
		ReferencesOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.TickingArea.Definition.Get()
		},
	},
	mcfunction.TagTypeFamilyId: {
		Kind:    References,
		Binding: stores.EntityFamily,
	},
	mcfunction.TagScoreboardObjectiveId: {
		Kind:    Definition,
		Binding: stores.ScoreboardObjective,
		// ReferencesOverride already returns Binding.Definition; the default would duplicate rename edits
		DefinitionOverride: func(node mcfunction.Node) []core.Symbol {
			return nil
		},
		ReferencesOverride: func(node mcfunction.Node) []core.Symbol {
			return stores.ScoreboardObjective.Definition.Get()
		},
	},
}

var rawMessage = &JsonRawMessageHandler{
	JsonSnippetHandler: &JsonSnippetHandler{
		JsonHandler: &JsonHandler{
			Entries: []JsonEntry{
				{
					Kind:    References,
					Binding: stores.Lang,
					Path:    []shared.JsonPath{shared.JsonValue("rawtext/**/translate")},
				},
				{
					Kind:    Definition,
					Binding: stores.ScoreboardObjective,
					Path:    []shared.JsonPath{shared.JsonValue("rawtext/**/score/objective")},
					DefinitionOverride: func(ctx *JsonContext) []core.Symbol {
						return nil
					},
					ReferencesOverride: func(ctx *JsonContext) []core.Symbol {
						return stores.ScoreboardObjective.Definition.Get()
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
				Kind:    References,
				Binding: stores.ItemId,
				Path: []shared.JsonPath{
					shared.JsonValue("minecraft:can_place_on/blocks/*"),
					shared.JsonValue("minecraft:can_destroy/blocks/*"),
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
