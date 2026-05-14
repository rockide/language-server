package lexer

import (
	"iter"
)

type Lexer struct {
	src           []rune
	i             int
	escapedQuotes bool
	isNewline     bool
}

func New(input []rune) *Lexer {
	return &Lexer{
		src:           input,
		escapedQuotes: false,
		isNewline:     true,
	}
}

func (l *Lexer) SetEscapedQuotes(value bool) {
	l.escapedQuotes = value
}

func (l *Lexer) eof() bool {
	return l.i >= len(l.src)
}

func (l *Lexer) peek() rune {
	if l.eof() {
		return 0
	}
	return l.src[l.i]
}

func (l *Lexer) peekN(n int) rune {
	if l.i+n >= len(l.src) {
		return 0
	}
	return l.src[l.i+n]
}

func (l *Lexer) previous() rune {
	if l.i-1 < 0 {
		return 0
	}
	return l.src[l.i-1]
}

func (l *Lexer) advance() rune {
	if l.eof() {
		return 0
	}
	r := l.src[l.i]
	l.i++
	return r
}

func (l *Lexer) advanceWhile(cond func(rune) bool) {
	for !l.eof() && cond(l.peek()) {
		l.advance()
	}
}

func (l *Lexer) pos() uint32 {
	return uint32(l.i)
}

func (l *Lexer) emit(kind TokenKind, start uint32) Token {
	end := l.pos()
	if kind == TokenNewline {
		l.isNewline = true
	} else {
		l.isNewline = false
	}
	return Token{
		Kind:  kind,
		Start: start,
		End:   end,
	}
}

func (l *Lexer) matchPairs(open, close rune) bool {
	depth := 0
	for !l.eof() {
		r := l.peek()
		if r == open {
			depth++
		} else if r == close {
			depth--
			if depth == 0 {
				l.advance()
				return true
			}
		} else if isNewline(r) {
			return false
		}
		l.advance()
	}
	return false
}

func (l *Lexer) Next() iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for !l.eof() {
			r := l.peek()
			start := l.pos()

			switch r {
			case '\r':
				l.advance()
				continue
			case '=':
				l.advance()
				if !yield(l.emit(TokenEquals, start)) {
					return
				}
				continue
			case ',':
				l.advance()
				if !yield(l.emit(TokenComma, start)) {
					return
				}
			case '!':
				l.advance()
				if !yield(l.emit(TokenBang, start)) {
					return
				}
				continue
			case '.':
				l.advance()
				if l.peek() == '.' {
					l.advance()
					if !yield(l.emit(TokenRange, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenDot, start)) {
						return
					}
				}
				continue
			case '-':
				if isDigit(l.peekN(1)) {
					l.advance() // consume '-'
					l.scanNumber()
					if !yield(l.emit(TokenNumber, start)) {
						return
					}
					continue
				}
			case '#':
				if l.isNewline && l.peekN(1) == '#' {
					l.advanceWhile(func(r rune) bool { return !isNewline(r) })
					if !yield(l.emit(TokenComment, start)) {
						return
					}
					continue
				}
			case '@':
				l.advance() // consume '@'
				a := l.i
				l.advanceWhile(isLetter)
				b := l.i
				if a == b {
					if !yield(l.emit(TokenString, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenSelector, start)) {
						return
					}
				}
				continue
			case '[':
				if l.matchPairs('[', ']') {
					if !yield(l.emit(TokenMap, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenUnterminatedMap, start)) {
						return
					}
				}
				continue
			case '{':
				if l.matchPairs('{', '}') {
					if !yield(l.emit(TokenJSON, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenUnterminatedJSON, start)) {
						return
					}
				}
				continue
			case '~', '^':
				l.advance() // consume '~' or '^'
				r = l.peek()
				if r == '-' {
					l.advance() // consume '-'
					r = l.peek()
				}
				if r == '.' {
					l.advance() // consume '.'
					l.advanceWhile(isDigit)
				} else {
					l.scanNumber()
				}
				if !yield(l.emit(TokenRelativeNumber, start)) {
					return
				}
				continue
			case '"':
				l.scanString()
				if l.previous() == '"' {
					if !yield(l.emit(TokenString, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenUnterminatedString, start)) {
						return
					}
				}
				continue
			case '\\':
				if l.escapedQuotes && l.peekN(1) == '"' {
					l.scanString()
					if l.previous() == '"' {
						if !yield(l.emit(TokenString, start)) {
							return
						}
					} else {
						if !yield(l.emit(TokenUnterminatedString, start)) {
							return
						}
					}
					continue
				}
			}

			if isWhitespace(r) {
				l.advanceWhile(isWhitespace)
				if !yield(l.emit(TokenWhitespace, start)) {
					return
				}
				continue
			}

			if isNewline(r) {
				l.advance()
				if !yield(l.emit(TokenNewline, start)) {
					return
				}
				continue
			}

			if isDigit(r) {
				l.scanNumber()
				if isIdent(l.peek()) {
					l.advanceWhile(isIdent)
					if !yield(l.emit(TokenString, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenNumber, start)) {
						return
					}
				}
				continue
			}

			l.advanceWhile(func(r rune) bool { return !isWhitespace(r) && !isNewline(r) })
			if !yield(l.emit(TokenString, start)) {
				return
			}

		}
	}
}

func (l *Lexer) Reset(input []rune) {
	l.src = input
	l.i = 0
	l.isNewline = true
}

func (l *Lexer) scanNumber() {
	l.advanceWhile(isDigit)
	if l.peek() == '.' && isDigit(l.peekN(1)) {
		l.advance() // consume '.'
		l.advanceWhile(isDigit)
	}
}

func (l *Lexer) scanString() {
	if l.escapedQuotes {
		l.advance() // consume opening '\'
	}
	l.advance() // consume opening '"'
	l.advanceWhile(func(r rune) bool { return r != '"' && !isNewline(r) })
	if l.peek() == '"' {
		l.advance()
	}
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isNewline(r rune) bool {
	return r == '\n'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isIdent(r rune) bool {
	return isLetter(r) || r == '_' || r == '-' || r == ':' || r == '§'
}
