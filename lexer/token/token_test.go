package token

import "testing"

func TestLookupIdentifier(t *testing.T) {
	tokenType := LookupIdentfier("if")
	if tokenType != IF {
		t.Errorf("tokenType is not of expected type. expected=%s, got=%s", IF, tokenType)
	}
}
