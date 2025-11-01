package lexer

import "unicode/utf8"

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

//readCharacter -> read the next character position

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

//peekCharacter returns next char

func (lexer *Lexer) peekCharacter() rune {
	if lexer.readPostion >= len(lexer.input) {
		return 0
	}
	ch, _ := utf8.DecodeLastRuneInString(lexer.input[lexer.readPostion:])
	return ch
}
