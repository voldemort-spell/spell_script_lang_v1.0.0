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

	// keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	FOR      TokenType = "FOR"
	WHILE    TokenType = "WHILE"
)

// Token -> reperesent lexical token
type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}
