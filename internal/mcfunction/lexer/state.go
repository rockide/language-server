package lexer

type state uint8

const (
	StateStart state = iota
	StateString
	StateNumber
	StateNegativeNumber
	StateRelativeNumber
	StateSelector
	StateMap
	StateJSON
)
