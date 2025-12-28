package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Structure = &Path{
	Pattern: shared.StructureGlob,
	Store:   stores.StructurePath,
}
