package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Texture = &Path{
	Pattern: shared.TextureGlob,
	Store:   stores.TexturePath,
}
