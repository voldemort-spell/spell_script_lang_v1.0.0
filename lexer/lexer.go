package lexer

import (
	"unicode"
	"unicode/utf8"
)

// Lexer ----> lexical analyzer

type Lexer struct {
	input       string
	position    int  // current position input
	readPostion int  //curr reading position
	ch          rune //curr char under examin
	pos         Position
}

// create new lexer

func NewLexer(input string) *Lexer {
	lex := &Lexer{input: input}
	return lex
}

// readCharacter -> read the next character position

func (lexer *Lexer) readCharacter() {
	if lexer.readPostion >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch, _ = utf8.DecodeLastRuneInString(lexer.input[lexer.readPostion:])
	}

	lexer.position = lexer.readPostion
	lexer.readPostion += utf8.RuneLen(lexer.ch)
	lexer.pos.Offset = lexer.position
	lexer.pos.Column++

	if lexer.ch == '\n' {
		lexer.pos.Line++
		lexer.pos.Column = 0
	}
}

// peekCharacter returns next char

func (lexer *Lexer) peekCharacter() rune {
	if lexer.readPostion >= len(lexer.input) {
		return 0
	}
	ch, _ := utf8.DecodeLastRuneInString(lexer.input[lexer.readPostion:])
	return ch
}

// check character is letter
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// check cracter is digit
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}

// read identifier from input
func (lexer *Lexer) readIdentfi() string {
	position := lexer.position

	for isLetter(lexer.ch) || isDigit(lexer.ch) {
		lexer.readCharacter()
	}

	return lexer.input[position:lexer.position]
}

// read number in input
func (lexer *Lexer) readNum() string {
	position := lexer.position

	for isDigit(lexer.ch) {
		lexer.readCharacter()
	}

	// decimal

	if lexer.ch == '.' && isDigit(lexer.peekCharacter()) {
		lexer.readCharacter()
		for isDigit(lexer.ch) {
			lexer.readCharacter()
		}
	}

	return lexer.input[position:lexer.position]
}

// return the next token in the input

func (lexer *Lexer) NextToken() Token {
	var token Token

	// skip witespace

	switch lexer.ch {
	case '=':
		token = Token{Type: EQUAL, Literal: string(lexer.ch)}
	case '+':
		token = Token{Type: PLUS, Literal: string(lexer.ch)}
	case '-':
		token = Token{Type: MINUS, Literal: string(lexer.ch)}
	case '!':
		token = Token{Type: BANG, Literal: string(lexer.ch)}
	case '/':
		token = Token{Type: SLASH, Literal: string(lexer.ch)}
	case '*':
		token = Token{Type: ASTERISK, Literal: string(lexer.ch)}
	case '<':
		token = Token{Type: GRATERTHAN, Literal: string(lexer.ch)}
	case '>':
		token = Token{Type: LESSTHAN, Literal: string(lexer.ch)}
	case ';':
		token = Token{Type: SEMICO, Literal: string(lexer.ch)}
	case ':':
		token = Token{Type: COLON, Literal: string(lexer.ch)}
	case ',':
		token = Token{Type: COMMA, Literal: string(lexer.ch)}
	case '(':
		token = Token{Type: LPAREN, Literal: string(lexer.ch)}
	case ')':
		token = Token{Type: RPAREN, Literal: string(lexer.ch)}
	case '{':
		token = Token{Type: LBRACE, Literal: string(lexer.ch)}
	case '}':
		token = Token{Type: RBRACE, Literal: string(lexer.ch)}
	case '[':
		token = Token{Type: LBRACKET, Literal: string(lexer.ch)}
	case ']':
		token = Token{Type: RBRACKET, Literal: string(lexer.ch)}
	case '"':
		token.Type = STR
		//token.Literal = l.re
	}
	return token
}
