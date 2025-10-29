package lexer

import "fmt"

type Position struct {
	Line   int
	Column int
	Offset int
}

func (p Position) String() string {
	return fmt.Sprintf("%d%d", p.Line, p.Column)
}
