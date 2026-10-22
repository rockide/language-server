package core

import (
	"path/filepath"

	"go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type Project struct {
	BP string
	RP string
}

func NewProject(bp string, rp string) *Project {
	project := Project{}
	if bp != "" {
		project.BP = filepath.ToSlash(filepath.Clean(bp))
	}
	if rp != "" {
		project.RP = filepath.ToSlash(filepath.Clean(rp))
	}
	return &project
}

type Symbol struct {
	Value string
	URI   uri.URI
	Range *protocol.Range
}
