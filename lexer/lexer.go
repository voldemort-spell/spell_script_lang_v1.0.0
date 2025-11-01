package lexer

// Lexer ----> lexical analyzer

type Lexer struct {
	input       string
	position    int  // current position input
	readPostion int  //curr reading position
	ch          rune //curr char under examin
	pos         Position
}

// create new lexer

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
