package lexer

type TokenKind uint8

const (
	TokenUnknown TokenKind = iota
	TokenWhitespace
	TokenNewline
	TokenComment
	TokenString
	TokenUnterminatedString
	TokenNumber
	TokenRelativeNumber
	TokenSelector
	TokenMap
	TokenUnterminatedMap
	TokenJSON
	TokenUnterminatedJSON
	TokenEquals
	TokenComma
	TokenDot
	TokenRange
	TokenBang
)

type Token struct {
	Kind  TokenKind
	Start uint32
	End   uint32
}

func (t Token) Text(input []rune) string {
	return string(input[t.Start:t.End])
}

func (t Token) Range() (start, end uint32) {
	return t.Start, t.End
}

func (t Token) Len() uint32 {
	return t.End - t.Start
}

func (t Token) IsInside(pos uint32) bool {
	return t.Start <= pos && pos < t.End
}
