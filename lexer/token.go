package lexer

// token type represents the type of token
type TokenType string


const (
	// specials tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF TokenType = "EOF"

	// identifiers and literals
	IDENT TokenType = "IDENT"
	INT TokenType = "INT"

	
)