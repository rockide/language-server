package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rockide/language-server/core"
	"github.com/rockide/language-server/internal/semtok"
	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	if params.ProcessID != nil {
		log.Printf("Process ID: %d", *params.ProcessID)
	}
	clientVersion, _ := params.ClientInfo.Version.Get()
	log.Printf("Connected to: %s %s", params.ClientInfo.Name, clientVersion)

	var options *core.RockideOptions
	if len(params.InitializationOptions) > 0 {
		options = &core.RockideOptions{}
		if err := json.Unmarshal(params.InitializationOptions, options); err != nil {
			log.Println(err)
			options = nil
		}
	}
	if err := s.workspace.LoadPacks(options); err != nil {
		log.Println(err)
		return &protocol.InitializeResult{}, nil
	}

	yes := true
	triggerCharacters := strings.Split(`0123456789abcdefghijklmnopqrstuvwxyz.'"() `, "")
	result := protocol.InitializeResult{
		ServerInfo: protocol.ServerInfo{
			Name:    "rockide",
			Version: protocol.NewOptional(s.version),
		},
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncKindIncremental,
			CompletionProvider: &protocol.CompletionOptions{
				TriggerCharacters: triggerCharacters,
			},
			DefinitionProvider: protocol.Boolean(true),
			ReferencesProvider: protocol.Boolean(true),
			RenameProvider: &protocol.RenameOptions{
				PrepareProvider: &yes,
			},
			HoverProvider: protocol.Boolean(true),
			SignatureHelpProvider: &protocol.SignatureHelpOptions{
				TriggerCharacters: triggerCharacters,
			},
			SemanticTokensProvider: &protocol.SemanticTokensOptions{
				Legend: protocol.SemanticTokensLegend{
					TokenTypes:     semtok.Types,
					TokenModifiers: semtok.Modifiers,
				},
				Full: protocol.Boolean(true),
			},
		},
	}
	return &result, nil
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	client, ok := protocol.ClientFromContext(ctx)
	if !ok {
		return nil
	}

	bp, rp := s.workspace.BehaviorPack, s.workspace.ResourcePack
	if bp == "" && rp == "" {
		return nil
	}

	var watcherPattern string
	if bp != "" && rp != "" {
		watcherPattern = fmt.Sprintf("{%s,%s}/**/*", bp, rp)
	} else if bp != "" {
		watcherPattern = bp + "/**/*"
	} else {
		watcherPattern = rp + "/**/*"
	}
	registerOptions, err := protocol.Marshal(protocol.DidChangeWatchedFilesRegistrationOptions{
		Watchers: []protocol.FileSystemWatcher{{
			GlobPattern: &protocol.RelativePattern{
				BaseURI: protocol.URI(uri.File(s.workspace.RootDir)),
				Pattern: protocol.Pattern(watcherPattern),
			},
		}},
	})
	if err != nil {
		return err
	}
	fileWatcher := protocol.Registration{
		ID:              "fileWatcher",
		Method:          "workspace/didChangeWatchedFiles",
		RegisterOptions: registerOptions,
	}
	if err := client.RegisterCapability(ctx, &protocol.RegistrationParams{Registrations: []protocol.Registration{fileWatcher}}); err != nil {
		return err
	}

	token := protocol.String(fmt.Sprintf("indexing-workspace-%d", time.Now().Unix()))
	if err := client.WorkDoneProgressCreate(ctx, &protocol.WorkDoneProgressCreateParams{Token: token}); err != nil {
		return err
	}
	progress := protocol.ProgressParams{Token: token}
	progress.Value, err = protocol.Marshal(&protocol.WorkDoneProgressBegin{Kind: "begin", Title: "Rockide: Indexing workspace"})
	if err != nil {
		return err
	}
	if err := client.Progress(ctx, &progress); err != nil {
		return err
	}

	s.indexWorkspace()

	progress.Value, err = protocol.Marshal(&protocol.WorkDoneProgressEnd{Kind: "end"})
	if err != nil {
		return err
	}
	return client.Progress(ctx, &progress)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return nil
}

func (s *Server) DidOpen(ctx context.Context, params *protocol.DidOpenTextDocumentParams) error {
	s.workspace.Documents.Open(params.TextDocument.URI, params.TextDocument.Text)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	s.workspace.Documents.SyncIncremental(params.TextDocument.URI, params.ContentChanges)
	s.debouncer.Debounce(params.TextDocument.URI, func() {
		s.onChange(params.TextDocument.URI)
	})
	return nil
}

func (s *Server) DidSave(ctx context.Context, params *protocol.DidSaveTextDocumentParams) error {
	s.workspace.Documents.SyncFull(params.TextDocument.URI, params.Text)
	return nil
}

func (s *Server) DidClose(ctx context.Context, params *protocol.DidCloseTextDocumentParams) error {
	s.workspace.Documents.Close(params.TextDocument.URI)
	return nil
}

func (s *Server) Completion(ctx context.Context, params *protocol.CompletionParams) (protocol.CompletionResult, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.CompletionProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return protocol.CompletionItemSlice(handler.Completions(document, params.Position)), nil
}

func (s *Server) Definition(ctx context.Context, params *protocol.DefinitionParams) (protocol.DefinitionResult, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.DefinitionProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	links := handler.Definitions(document, params.Position)
	result := make(protocol.DefinitionLinkSlice, len(links))
	for i, link := range links {
		result[i] = protocol.DefinitionLink(link)
	}
	return result, nil
}

func (s *Server) References(ctx context.Context, params *protocol.ReferenceParams) ([]protocol.Location, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.ReferencesProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return handler.References(document, params.Position, params.Context.IncludeDeclaration), nil
}

func (s *Server) PrepareRename(ctx context.Context, params *protocol.PrepareRenameParams) (protocol.PrepareRenameResult, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.RenameProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	placeholder := handler.PrepareRename(document, params.Position)
	if placeholder == nil {
		return nil, nil
	}
	return placeholder, nil
}

func (s *Server) Rename(ctx context.Context, params *protocol.RenameParams) (*protocol.WorkspaceEdit, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.RenameProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return handler.Rename(document, params.Position, params.NewName), nil
}

func (s *Server) Hover(ctx context.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.HoverProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return handler.Hover(document, params.Position), nil
}

func (s *Server) SignatureHelp(ctx context.Context, params *protocol.SignatureHelpParams) (*protocol.SignatureHelp, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.SignatureHelpProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return handler.SignatureHelp(document, params.Position), nil
}

func (s *Server) SemanticTokensFull(ctx context.Context, params *protocol.SemanticTokensParams) (*protocol.SemanticTokens, error) {
	handler, ok := s.findHandler(params.TextDocument.URI).(core.SemanticTokenProvider)
	if !ok {
		return nil, nil
	}
	document := s.workspace.Documents.Get(params.TextDocument.URI)
	return handler.SemanticTokens(document), nil
}

func (s *Server) DidChangeWatchedFiles(ctx context.Context, params *protocol.DidChangeWatchedFilesParams) error {
	for _, change := range params.Changes {
		switch change.Type {
		case protocol.FileChangeTypeCreated:
			s.onCreate(change.URI)
		case protocol.FileChangeTypeChanged:
			s.debouncer.Debounce(change.URI, func() {
				s.onChange(change.URI)
			})
		case protocol.FileChangeTypeDeleted:
			s.debouncer.Cancel(change.URI)
			s.onDelete(change.URI)
		}
	}
	return nil
}
