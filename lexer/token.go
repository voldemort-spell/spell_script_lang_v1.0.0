package lexer

// token type represents the type of token
type TokenType string

const (
	// specials tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifiers and literals
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"
	STR   TokenType = "STR"

	// operators
	ASSIGN     TokenType = "="
	PLUS       TokenType = "+"
	MINUS      TokenType = "-"
	BAND       TokenType = "!"
	ASTERISK   TokenType = "*"
	SLASH      TokenType = "/"
	LESSTHAN   TokenType = "<"
	GRATERTHAN TokenType = ">"
	EQUAL      TokenType = "=="
	NOT_EQUAL  TokenType = "!="

	// delimiter
	COMMA  TokenType = "."
	SEMICO TokenType = ";"
	COLON  TokenType = ":"

	LPAREN   TokenType = "("
	RPAREN   TokenType = "}"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"
)
