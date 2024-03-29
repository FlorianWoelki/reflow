package ast

import (
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode() // Not used in production. Only used for not confusing the Go compiler.
}

type Expression interface {
	Node
	expressionNode() // Not used in production. Only used for not confusing the Go compiler.
}

type Program struct {
	Statements []Statement
}

func (p Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
