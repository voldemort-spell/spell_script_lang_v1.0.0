package lexer

import (
	"testing"
)

func TestNextToken_repl(t *testing.T) {
	input := `==+()`

	test1 := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{EQUAL, "=="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
	}

	lex := NewLexer(input)

	for i, tt := range test1 {
		tk := lex.NextToken()

		if tk.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type err. expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal err. expected=%q, got=%q",
				i, tt.expectedLiteral, tk.Literal)
		}

	}

	// v value1 = 100
}

func TestNextToken_code(t *testing.T) {
	input1 := `
		v  value1 = 100
		v  value2 = -120

		v sumeation = fn(x, y){
			return x + y
		}

		print(sumation)
	`

	test2 := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{

		{VAR, "v"},
		{IDENT, "value1"},
		{ASSIGN, "="},
		{NUMBER, "100"},

		{VAR, "v"},
		{IDENT, "value2"},
		{ASSIGN, "="},
		{MINUS, "-"},
		{NUMBER, "120"},

		{VAR, "v"},
		{IDENT, "summeation"},
		{ASSIGN, "="},
		{FUNCTION, "fn"},
		{LPAREN, "("},
		{IDENT, "x"},
		{COMMA, ","},
		{IDENT, "y"},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RETURN, "return"},
		{IDENT, "x"},
		{PLUS, "+"},
		{IDENT, "y"},
		{LBRACE, "}"},

		{PRINT, "print"},
		{LPAREN, "("},
		{IDENT, "summation"},
		{LPAREN, ")"},

		{EOF, ""},
	}

	lex1 := NewLexer(input1)

	for i, tt := range test2 {
		tok := lex1.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type err. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
