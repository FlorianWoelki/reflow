package evaluator

import (
	"strconv"
	"testing"

	"github.com/florianwoelki/reflow/object"
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

func TestBuiltinLen(t *testing.T) {
	result := builtinLen()
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("result for empty args is of wrong type. expected=%s, got=%s", object.ERROR_OBJ, result.Type())
	}

	str := "Hello World"
	result = builtinLen(&object.String{Value: str})
	if result.Type() != object.INTEGER_OBJ {
		t.Errorf("result for string is of wrong type. expected=%s, got=%s", object.INTEGER_OBJ, result.Type())
	}

	parsedLen, err := strconv.Atoi(result.Inspect())
	if err != nil {
		t.Errorf("received error while parsing length of string. got=%s", err)
	}

	if parsedLen != len(str) {
		t.Errorf("result for string is of wrong length. expected=%d, got=%d", len(str), parsedLen)
	}

	arr := &object.Array{Elements: []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 1}}}
	result = builtinLen(arr)
	if result.Type() != object.INTEGER_OBJ {
		t.Errorf("result for array is of wrong type. expected=%s, got=%s", object.INTEGER_OBJ, result.Type())
	}

	parsedLen, err = strconv.Atoi(result.Inspect())
	if err != nil {
		t.Errorf("received error while parsing length of array. got=%s", err)
	}

	if parsedLen != len(arr.Elements) {
		t.Errorf("result for array is of wrong length. expected=%d, got=%d", len(str), parsedLen)
	}

}
