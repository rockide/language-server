package handlers

import (
	"github.com/rockide/language-server/shared"
	"github.com/rockide/language-server/stores"
)

var Sound = &Path{
	Pattern: shared.SoundGlob,
	Store:   stores.SoundPath,
}
