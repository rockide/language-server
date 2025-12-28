package core

import (
	"path/filepath"

	"github.com/rockide/language-server/internal/protocol"
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
	URI   protocol.DocumentURI
	Range *protocol.Range
}
