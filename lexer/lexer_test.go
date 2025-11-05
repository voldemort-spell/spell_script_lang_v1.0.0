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
			t.Fatalf("tests[%d] - token type err. expected=%q, got=%q",
				i, tt.expectedType, tk.Type)
		}

		if tk.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal err. expected=%q, got=%q",
				i, tt.expectedLiteral, tk.Literal)
		}

	}

	input1 = `
		v  value1 = 100
		v  value2 = -120

		v sumeation = fn(x, y){
			return x + y
		}

		print(sumation)
	`
	lex1 := NewLexer(input1)

	tk1 := lex1.NextToken()
	if tk1.Type != VAR {
		t.Fatalf("expected VAR, got %q", tk1.Type)
	}
	if tk1.Literal != "v" {		t.Fatalf("expected 'v', got %q", tk1.Literal)
	}

	tk2 := lex1.NextToken()
	if tk2.Type != IDENT {
		t.Fatalf("expected IDENT, got %q", tk2.Type)
	}
	if tk2.Literal != "value1" {
		t.Fatalf("expected 'value1', got %q", tk2.Literal)
	}
	tk3 := lex1.NextToken()
	if tk3.Type != ASSIGN {
		t.Fatalf("expected ASSIGN, got %q", tk3.Type)
	}
	if tk3.Literal != "=" {
		t.Fatalf("expected '=', got %q", tk3.Literal)
	}

	tk4 := lex1.NextToken()
	if tk4.Type != INT {
		t.Fatalf("expected INT, got %q", tk4.Type)
	}
	if tk4.Literal != "100" {
		t.Fatalf("expected '100', got %q", tk4.Literal)
	}
	tk5 := lex1.NextToken()
	if tk5.Type != SEMICOLON {
		t.Fatalf("expected SEMICOLON, got %q", tk5.Type)
	// v value1 = 100

}
