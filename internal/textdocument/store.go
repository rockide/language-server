package textdocument

import (
	"os"
	"slices"
	"sync"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type Store struct {
	documents map[uri.URI]*TextDocument
	mu        sync.RWMutex
}

func NewStore() *Store {
	return &Store{documents: make(map[uri.URI]*TextDocument)}
}

func (s *Store) Get(uri uri.URI) *TextDocument {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.documents[uri]
}

func (s *Store) Open(uri uri.URI, txt string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	document := TextDocument{URI: uri, content: []rune(txt)}
	s.documents[uri] = &document
}

func (s *Store) Close(uri uri.URI) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.documents, uri)
}

func (s *Store) SyncIncremental(uri uri.URI, contentChanges []protocol.TextDocumentContentChangeEvent) {
	if len(contentChanges) == 0 {
		return
	}
	document := s.Get(uri)
	if document == nil {
		return
	}
	for _, change := range contentChanges {
		switch change := change.(type) {
		case *protocol.TextDocumentContentChangeWholeDocument:
			document.content = []rune(change.Text)
		case *protocol.TextDocumentContentChangePartial:
			startOffset := document.OffsetAt(change.Range.Start)
			endOffset := document.OffsetAt(change.Range.End)
			document.content = slices.Concat(document.content[:startOffset], []rune(change.Text), document.content[endOffset:])
		}
		document.lineOffsets = nil
	}
}

func (s *Store) SyncFull(uri uri.URI, txt *string) {
	if txt == nil {
		return
	}
	document := s.Get(uri)
	if document == nil || document.GetText() == *txt {
		return
	}
	document.content = []rune(*txt)
	document.lineOffsets = nil
}

func ReadFile(uri uri.URI) (*TextDocument, error) {
	b, err := os.ReadFile(uri.FsPath())
	if err != nil {
		return nil, err
	}
	document := TextDocument{URI: uri, content: []rune(string(b))}
	return &document, nil
}

func (s *Store) GetOrReadFile(uri uri.URI) (*TextDocument, error) {
	if document := s.Get(uri); document != nil {
		return document, nil
	}
	return ReadFile(uri)
}
