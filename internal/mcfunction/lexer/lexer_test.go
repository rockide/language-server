package lexer_test

import (
	"reflect"
	"testing"

	"github.com/rockide/language-server/internal/mcfunction/lexer"
)

func assertTokens(t *testing.T, input string, expected []lexer.Token) {
	parser := lexer.New([]rune(input))
	i := 0
	for token := range parser.Next() {
		if i >= len(expected) {
			t.Errorf("Unexpected token: %v", token)
			continue
		}
		if !reflect.DeepEqual(token, expected[i]) {
			t.Errorf("Input: %s, Expected: %v, Actual: %v", input, expected[i], token)
		}
		i++
	}
}

func assertEscapedTokens(t *testing.T, input string, expected []lexer.Token) {
	lex := lexer.New([]rune(input))
	lex.SetEscapedQuotes(true)
	i := 0
	for token := range lex.Next() {
		if i >= len(expected) {
			t.Errorf("Unexpected token: %v", token)
			continue
		}
		if !reflect.DeepEqual(token, expected[i]) {
			t.Errorf("Input: %s, Expected: %v, Actual: %v", input, expected[i], token)
		}
		i++
	}
}

func TestLexer(t *testing.T) {
	assertTokens(t, "9", []lexer.Token{
		{lexer.TokenNumber, 0, 1},
	})
	assertTokens(t, "90_degrees", []lexer.Token{
		{lexer.TokenString, 0, 10},
	})
	assertTokens(t, "90 degrees", []lexer.Token{
		{lexer.TokenNumber, 0, 2},
		{lexer.TokenWhitespace, 2, 3},
		{lexer.TokenString, 3, 10},
	})

	assertTokens(t, "~~~", []lexer.Token{
		{lexer.TokenRelativeNumber, 0, 1},
		{lexer.TokenRelativeNumber, 1, 2},
		{lexer.TokenRelativeNumber, 2, 3},
	})
	assertTokens(t, "~-2.4~.2~3", []lexer.Token{
		{lexer.TokenRelativeNumber, 0, 5},
		{lexer.TokenRelativeNumber, 5, 8},
		{lexer.TokenRelativeNumber, 8, 10},
	})
	assertTokens(t, "-24 44", []lexer.Token{
		{lexer.TokenNumber, 0, 3},
		{lexer.TokenWhitespace, 3, 4},
		{lexer.TokenNumber, 4, 6},
	})

	assertTokens(t, "[[[]]]", []lexer.Token{
		{lexer.TokenMap, 0, 6},
	})
	assertTokens(t, "[[[", []lexer.Token{
		{lexer.TokenUnterminatedMap, 0, 3},
	})

	assertTokens(t, `"foo bar"`, []lexer.Token{
		{lexer.TokenString, 0, 9},
	})
	assertTokens(t, `"foo bar`, []lexer.Token{
		{lexer.TokenUnterminatedString, 0, 8},
	})
	assertEscapedTokens(t, `\"foo bar\"`, []lexer.Token{
		{lexer.TokenString, 0, 11},
	})
	assertTokens(t, "---", []lexer.Token{
		{lexer.TokenString, 0, 3},
	})
	assertTokens(t, "§afoobar", []lexer.Token{
		{lexer.TokenString, 0, 8},
	})

	assertTokens(t, "@", []lexer.Token{
		{lexer.TokenString, 0, 1},
	})

	assertTokens(t, "execute as @s[x=4]", []lexer.Token{
		{lexer.TokenString, 0, 7},
		{lexer.TokenWhitespace, 7, 8},
		{lexer.TokenString, 8, 10},
		{lexer.TokenWhitespace, 10, 11},
		{lexer.TokenSelector, 11, 13},
		{lexer.TokenMap, 13, 18},
	})
	assertTokens(t, "say ## not a comment", []lexer.Token{
		{lexer.TokenString, 0, 3},
		{lexer.TokenWhitespace, 3, 4},
		{lexer.TokenString, 4, 6},
		{lexer.TokenWhitespace, 6, 7},
		{lexer.TokenString, 7, 10},
		{lexer.TokenWhitespace, 10, 11},
		{lexer.TokenString, 11, 12},
		{lexer.TokenWhitespace, 12, 13},
		{lexer.TokenString, 13, 20},
	})

	assertTokens(t, "/say hello", []lexer.Token{
		{lexer.TokenString, 0, 4},
		{lexer.TokenWhitespace, 4, 5},
		{lexer.TokenString, 5, 10},
	})
}
