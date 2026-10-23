package textdocument

import (
	"slices"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type TextDocument struct {
	URI         uri.URI `json:"uri"`
	content     []rune
	lineOffsets []uint32
}

func (d *TextDocument) getLineOffsets() []uint32 {
	if d.lineOffsets == nil {
		d.lineOffsets = computeLineOffsets(d.content, true, 0)
	}
	return d.lineOffsets
}

func (d *TextDocument) PositionAt(offset uint32) protocol.Position {
	contentLength := uint32(len(d.content))
	offset = min(offset, contentLength)
	lineOffsets := d.getLineOffsets()
	low := 0
	high := len(lineOffsets)
	for low < high {
		mid := (low + high) / 2
		if lineOffsets[mid] > offset {
			high = mid
		} else {
			low = mid + 1
		}
	}

	line := max(low-1, 0)
	lineStart := lineOffsets[line]
	runeCol := offset - lineStart

	var lineEnd uint32
	if line+1 < len(lineOffsets) {
		lineEnd = lineOffsets[line+1]
	} else {
		lineEnd = contentLength
	}

	lineRunes := d.content[lineStart:lineEnd]
	utf16Char := utf16Len(lineRunes[:runeCol])

	return protocol.Position{
		Line:      uint32(line),
		Character: utf16Char,
	}
}

func (d *TextDocument) OffsetAt(position protocol.Position) uint32 {
	lineOffsets := d.getLineOffsets()
	maxLine := uint32(len(lineOffsets))
	contentLength := uint32(len(d.content))
	if position.Line >= maxLine {
		return contentLength
	}

	lineStart := lineOffsets[position.Line]
	var lineEnd uint32
	if position.Line+1 < maxLine {
		lineEnd = lineOffsets[position.Line+1]
	} else {
		lineEnd = contentLength
	}

	lineRunes := d.content[lineStart:lineEnd]
	runeCol := utf16ToRuneOffset(lineRunes, position.Character)

	return lineStart + runeCol
}

func (d *TextDocument) GetText() string {
	return string(d.content)
}

func (d *TextDocument) GetContent() []rune {
	return d.content
}

func (d *TextDocument) CreateVirtualDocument(ranges ...protocol.Range) *TextDocument {
	content := slices.Clone(d.content)
	textLength := uint32(len(content))
	result := make([]rune, textLength)

	offsets := make([][2]uint32, len(ranges))
	for i, r := range ranges {
		offsets[i][0] = d.OffsetAt(r.Start)
		offsets[i][1] = d.OffsetAt(r.End)
	}

	for i := range textLength {
		ch := content[i]
		if isEOL(ch) {
			result[i] = ch
			continue
		}
		isInside := false
		for _, offset := range offsets {
			if i >= offset[0] && i < offset[1] {
				isInside = true
				break
			}
		}
		if isInside {
			result[i] = ch
		} else {
			result[i] = ' '
		}
	}

	return &TextDocument{
		URI:     d.URI,
		content: result,
	}
}
