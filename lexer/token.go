package lexer

// token type represents the type of token
type TokenType string

const (
	// specials tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// identifiers and literals
	IDENT  TokenType = "IDENT"
	NUMBER TokenType = "NUMBER"
	STR    TokenType = "STR"

	// operators
	ASSIGN     TokenType = "="
	PLUS       TokenType = "+"
	MINUS      TokenType = "-"
	BANG       TokenType = "!"
	ASTERISK   TokenType = "*"
	SLASH      TokenType = "/"
	LESSTHAN   TokenType = "<"
	GRATERTHAN TokenType = ">"
	EQUAL      TokenType = "=="
	NOT_EQUAL  TokenType = "!="

	// delimiter
	COMMA  TokenType = ","
	SEMICO TokenType = ";"
	COLON  TokenType = ":"

	LPAREN   TokenType = "("
	RPAREN   TokenType = ")"
	LBRACE   TokenType = "{"
	RBRACE   TokenType = "}"
	LBRACKET TokenType = "["
	RBRACKET TokenType = "]"

	// keywords

	// LET represent -> v (variable)
	// define a variable v <variable_name> = <value>

	FUNCTION TokenType = "FUNCTION"
	VAR      TokenType = "VAR"
	TRUE     TokenType = "TRUE"
	FALSE    TokenType = "FALSE"
	IF       TokenType = "IF"
	ELSE     TokenType = "ELSE"
	RETURN   TokenType = "RETURN"
	FOR      TokenType = "FOR"
	WHILE    TokenType = "WHILE"
	PRINT    TokenType = "PRINT"
	PRINTLN  TokenType = "PRINTLN"
	
)

// Token -> reperesent lexical token
type Token struct {
	Type    TokenType
	Literal string
	Pos     Position
}

// check keywords
func LookIdent(identtifer string) TokenType {

	
	keywords := map[string]TokenType{
		"fn":     FUNCTION,
		"v":      VAR,
		"true":   TRUE,
		"false":  FALSE,
		"if":     IF,
		"else":   ELSE,
		"return": RETURN,
		"for":    FOR,
		"while":  WHILE,
		"print":  PRINT,
	}

	if token, ok := keywords[identtifer]; ok {
		return token
	}

	return IDENT
}
