package lexer

import (
	Token "LangLang/internals/token"
	"testing"
)

// symbols
func TestNextToken_single_symbols(t *testing.T) {
	input := `=+(){},`

	tests := []struct {
		expectedType    Token.TokenType
		expectedLiteral string
	}{
		{Token.ASSIGN, "="},
		{Token.PLUS, "+"},
		{Token.LPAREN, "("},
		{Token.RPAREN, ")"},
		{Token.LBRACE, "{"},
		{Token.RBRACE, "}"},
		{Token.COMMA, ","},
		{Token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}

	t.Log("Test 1 finishs")
}

func TestNextToken_expression(t *testing.T) {
	input := `
	x = 9
	y = 10
	print x + y
	`

	tests := []struct {
		expectedType    Token.TokenType
		expectedLiteral string
	}{
		{Token.IDENT, "x"},
		{Token.ASSIGN, "="},
		{Token.INT, "9"},
		{Token.IDENT, "y"},
		{Token.ASSIGN, "="},
		{Token.INT, "10"},
		{Token.IDENT, "print"},
		{Token.IDENT, "x"},
		{Token.PLUS, "+"},
		{Token.IDENT, "y"},
		{Token.EOF, ""},
	}

	l := NewLexer(input)
	t.Log("before loop")
	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, token.Literal)
		}
	}
	t.Log("Test 2 finishs")
}
