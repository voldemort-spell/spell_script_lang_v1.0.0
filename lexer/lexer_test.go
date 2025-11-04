package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `==+()`

	test := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{EQUAL, "=="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
	}

	lex := NewLexer(input)

	for i, tt := range test {
		tk := lex.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tk.Literal)
		}

	}
}
