package core

import "github.com/rockide/language-server/internal/protocol"

type Project struct {
	BP string
	RP string
}

type Symbol struct {
	Value string
	URI   protocol.DocumentURI
	Range *protocol.Range
}
