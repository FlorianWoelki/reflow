package ast

import (
	"bytes"

	"github.com/florianwoelki/reflow/lexer/token"
)

type IndexExpression struct {
	Token      token.Token
	Left       Expression
	Index      Expression
	Assignment Expression
}

func (ie *IndexExpression) expressionNode() {}

func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")

	if ie.Assignment != nil {
		out.WriteString(" = ")
		out.WriteString(ie.Assignment.String())
	}

	return out.String()
}
