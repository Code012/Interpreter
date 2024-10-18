package ast

import "monkey-go/token"

// Represents a single node in the ast
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode() // dummy method,  they are not strictly necessary but help us by guiding the Go compiler and possibly causing it to throw errors when we use a Statement where an Expression should’ve been used, and vice versa.
}

type Expression interface {
	Node
	expressionNode() //dummy method,  they are not strictly necessary but help us by guiding the Go compiler and possibly causing it to throw errors when we use a Statement where an Expression should’ve been used, and vice versa.
}

// A Monkey program is a series of statemetents, all those statements are contained in the Statements slice, which i sjust a slice of AST nodes that implement the Statement interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Represents identifiers even if they are not expressions (dont produce values) for simplicity and to keep the number of node types small. e.g. the x in let x = 5 is a statement because it does not produce a value however the x in let x = add(2,3) is an expression
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
