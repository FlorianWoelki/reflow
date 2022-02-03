package ast

import (
	"bytes"

	"github.com/florianwoelki/reflow/token"
)

type AssignmentStatement struct {
	Token      token.Token
	Operator   string
	Name       *Identifier
	Value      Expression
	Assignment Expression
}

func (as *AssignmentStatement) statementNode() {}

func (as *AssignmentStatement) TokenLiteral() string {
	return ""
}

func (as *AssignmentStatement) String() string {
	var out bytes.Buffer

	if as.Operator != token.ASSIGN {
		out.WriteString(as.Name.String())
		out.WriteString(" = ")
	}

	out.WriteString(as.Name.String())
	out.WriteString(" " + as.Operator + " ")

	out.WriteString(as.Value.String())

	out.WriteString(";")
	return out.String()
}
