package lexer

import (
	"iter"
)

type Lexer struct {
	src     []rune
	i       int
	state   state
	lastPos int
}

func New(input []rune) *Lexer {
	return &Lexer{
		src:     input,
		state:   StateStart,
		lastPos: -1,
	}
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

func (l *Lexer) advance() rune {
	if l.eof() {
		return 0
	}
	l.lastPos = l.i
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

func (l *Lexer) isInfiniteLoop() bool {
	return l.i == l.lastPos
}

func (l *Lexer) Next() iter.Seq[Token] {
	return func(yield func(Token) bool) {
		for !l.eof() {
			if l.isInfiniteLoop() {
				panic("possible infinite loop in lexer")
			}
			r := l.peek()
			start := l.pos()
			switch r {
			case '\r':
				l.advance()
				continue
			case ' ', '\t':
				l.advanceWhile(isWhitespace)
				if !yield(l.emit(TokenWhitespace, start)) {
					return
				}
				continue
			case '\n':
				l.advance()
				if !yield(l.emit(TokenNewline, start)) {
					return
				}
				l.state = StateStart
				continue
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				l.state = StateNumber
			case '-':
				if l.i+1 < len(l.src) && isDigit(l.src[l.i+1]) {
					l.state = StateNegativeNumber
				}
			case '~', '^':
				l.state = StateRelativeNumber
			case '@':
				l.state = StateSelector
			case '[':
				l.state = StateMap
			case '{':
				l.state = StateJSON
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
			case '!':
				l.advance()
				if !yield(l.emit(TokenBang, start)) {
					return
				}
				continue
			}

			switch l.state {
			case StateStart:
				if r == '#' {
					if l.peek() == '#' {
						l.advanceWhile(func(r rune) bool { return !isNewline(r) })
						if !yield(l.emit(TokenComment, start)) {
							return
						}
						continue
					}
				}
				fallthrough
			case StateString:
				if r == '"' {
					l.advance()
					l.advanceWhile(func(r rune) bool { return r != '"' && !isNewline(r) })
					if l.peek() == '"' {
						l.advance()
						if !yield(l.emit(TokenString, start)) {
							return
						}
					} else {
						if !yield(l.emit(TokenUnterminatedString, start)) {
							return
						}
						l.state = StateStart
						continue
					}
				} else {
					l.advanceWhile(func(r rune) bool { return !isTerminateString(r) })
					if !yield(l.emit(TokenString, start)) {
						return
					}
				}
			case StateNegativeNumber:
				l.advance() // consume '-'
				a := l.i
				l.advanceWhile(isDigit)
				b := l.i
				r := l.peek()
				if r == '.' && isDigit(l.peekN(1)) {
					l.advance()
					l.advanceWhile(isDigit)
				}
				if a == b {
					if !yield(l.emit(TokenString, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenNumber, start)) {
						return
					}
				}
			case StateNumber:
				l.advanceWhile(isDigit)
				r := l.peek()
				if r == '.' && isDigit(l.peekN(1)) {
					l.advance() // consume '.'
					l.advanceWhile(isDigit)
				}
				if !yield(l.emit(TokenNumber, start)) {
					return
				}
			case StateRelativeNumber:
				// possible patterns:
				// ~
				// ~2
				// ~2.2
				// ~.2
				// ~-2
				// ~-2.2
				// ~-.2
				l.advance() // consume '~' or '^'
				r := l.peek()
				if r == '-' {
					l.advance() // consume '-'
					r = l.peek()
				}
				if isDigit(r) {
					l.advanceWhile(isDigit)
					r = l.peek()
				}
				if r == '.' {
					l.advance() // consume '.'
					l.advanceWhile(isDigit)
				}
				if !yield(l.emit(TokenRelativeNumber, start)) {
					return
				}
			case StateSelector:
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
			case StateMap:
				if l.matchPairs('[', ']') {
					if !yield(l.emit(TokenMap, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenUnterminatedMap, start)) {
						return
					}
				}
			case StateJSON:
				if l.matchPairs('{', '}') {
					if !yield(l.emit(TokenJSON, start)) {
						return
					}
				} else {
					if !yield(l.emit(TokenUnterminatedJSON, start)) {
						return
					}
				}
			}
			l.state = StateString
		}
	}
}

func (l *Lexer) Reset(input []rune) {
	l.src = input
	l.i = 0
	l.lastPos = -1
	l.state = StateStart
}

func isWhitespace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\r'
}

func isNewline(r rune) bool {
	return r == '\n'
}

func isTerminateString(r rune) bool {
	return isNewline(r) || isWhitespace(r) || isRelativeNumber(r) || r == '@' || r == '"' || r == '[' || r == '{' || r == '=' || r == ',' || r == '!'
}

func isRelativeNumber(r rune) bool {
	return r == '~' || r == '^'
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
