package evaluator

import (
	"testing"
)

func TestBuiltins(t *testing.T) {
	expectedFunctions := []string{"len", "print", "str", "first", "last", "rest", "pop", "push", "map", "find", "filter"}

	if len(builtins) != len(expectedFunctions) {
		t.Errorf("length of expected functions is not the same as builtins. expected=%d, got=%d", len(builtins), len(expectedFunctions))
	}

	for _, function := range expectedFunctions {
		_, ok := builtins[function]
		if !ok {
			t.Errorf("function is not in builtins function map. expected=%s", function)
			continue
		}
	}
}
