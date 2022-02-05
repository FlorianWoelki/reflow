package ast

import (
	"bytes"

	"github.com/florianwoelki/reflow/lexer/token"
)

type PostfixExpression struct {
	Token    token.Token
	Name     Expression
	Operator string
}

func (pe *PostfixExpression) expressionNode() {}

func (pe *PostfixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PostfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Name.String())
	out.WriteString(pe.Operator)
	out.WriteString(")")

	return out.String()
}
