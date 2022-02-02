package ast

import (
	"testing"

	"github.com/florianwoelki/reflow/token"
)

func TestString(t *testing.T) {
	result := "let someTest = anotherTest;"
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "someTest"},
					Value: "someTest",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherTest"},
					Value: "anotherTest",
				},
			},
		},
	}

	if program.String() != result {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
