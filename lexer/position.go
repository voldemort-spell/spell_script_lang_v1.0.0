package lexer

import "fmt"

type Position struct {
	Line   int
	Column int
	Offset int
}

func (p Position) String() string {
	return fmt.Sprintf("line %d, col %d", p.Line, p.Column)
}
