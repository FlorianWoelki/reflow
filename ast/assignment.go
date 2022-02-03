package ast

import (
	"bytes"

	"github.com/florianwoelki/reflow/token"
)

type AssignmentStatement struct {
	Token    token.Token
	Operator string
	Name     *Identifier
	Value    Expression
}

func (as *AssignmentStatement) statementNode() {}

func (as *AssignmentStatement) TokenLiteral() string {
	return ""
}

func (as *AssignmentStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Name.String())
	out.WriteString(" " + as.Operator + " ")

	if as.Value != nil {
		out.WriteString(as.Value.String())
	}

	out.WriteString(";")
	return out.String()
}
