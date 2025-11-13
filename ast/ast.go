package ast

import "go/token"

// AST - abstract tree syntax
// v <identifier> = <expression>;

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	StatementNode()
}

type Expression interface {
	Node
	ExpressionNode()
}

type Programe struct {
	Statement []Statement
}

func (program *Programe) TokenLiteral() string {
	if len(program.Statement) > 0 {
		return program.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

type Ident struct {
	Token token.Token
	Value string
}

type VarStatement struct {
	Token token.Token
	Name  *Ident
}
