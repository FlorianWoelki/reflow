package evaluator

import (
	"strconv"
	"testing"

	"github.com/florianwoelki/reflow/object"
)

func TestBuiltinsRest(t *testing.T) {
	result := builtinRest()
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call without parameters gave no error back. got=%s", result.Type())
	}

	result = builtinRest(&object.Integer{Value: 3})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call with wrong parameter gave no error back. got=%s", result.Type())
	}

	elements := []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}, &object.Integer{Value: 3}, &object.Integer{Value: 4}}
	result = builtinRest(&object.Array{Elements: elements})
	if result.Type() == NULL.Type() {
		t.Errorf("function call with array parameter did not return NULL. got=%s", result.Type())
	}

	if result.Type() != object.ARRAY_OBJ {
		t.Errorf("function call with array parameter returned the wrong type. expected=%s, got=%s", object.ARRAY_OBJ, result.Type())
	}

	if result.Inspect() != "[2, 3, 4]" {
		t.Errorf("function call with array parameter returned the wrong value. expected=%s, got=%s", "[2, 3, 4]", result.Inspect())
	}
}

func TestBuiltinsLast(t *testing.T) {
	result := builtinLast()
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call without parameters gave no error back. got=%s", result.Type())
	}

	result = builtinLast(&object.Integer{Value: 3})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call with wrong parameter gave no error back. got=%s", result.Type())
	}

	elements := []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}}
	result = builtinLast(&object.Array{Elements: elements})
	if result.Type() == NULL.Type() {
		t.Errorf("function call with array parameter did not return NULL. got=%s", result.Type())
	}

	if result.Inspect() != elements[len(elements)-1].Inspect() {
		t.Errorf("function call with array parameter did not return the correct result. expected=%s, got=%s", elements[0].Inspect(), result.Inspect())
	}
}

func TestBuiltinsFirst(t *testing.T) {
	result := builtinFirst()
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call without parameters gave no error back. got=%s", result.Type())
	}

	result = builtinFirst(&object.Integer{Value: 3})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("function call with wrong parameter gave no error back. got=%s", result.Type())
	}

	elements := []object.Object{&object.Integer{Value: 1}, &object.Integer{Value: 2}}
	result = builtinFirst(&object.Array{Elements: elements})
	if result.Type() == NULL.Type() {
		t.Errorf("function call with array parameter did not return NULL. got=%s", result.Type())
	}

	if result.Inspect() != elements[0].Inspect() {
		t.Errorf("function call with array parameter did not return the correct result. expected=%s, got=%s", elements[0].Inspect(), result.Inspect())
	}
}

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

	result = builtinLen(&object.Integer{Value: 1})
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("result for integer should return an error. got=%v (%+v)", result, result)
	}
}

func TestBuiltinStr(t *testing.T) {
	tests := []struct {
		input    object.Object
		expected string
	}{
		{
			input:    &object.Array{Elements: []object.Object{&object.Integer{Value: 1}}},
			expected: "ERROR: argument to `str` cannot be ARRAY",
		},
		{
			input:    &object.Boolean{Value: true},
			expected: "true",
		},
		{
			input:    &object.Boolean{Value: false},
			expected: "false",
		},
		{
			input:    &object.Integer{Value: 1},
			expected: "1",
		},
		{
			input:    &object.String{Value: "hello"},
			expected: "hello",
		},
	}

	result := builtinStr()
	if result.Type() != object.ERROR_OBJ {
		t.Errorf("wrong result of type for empty args. expected=%s, got=%s", object.ERROR_OBJ, result.Type())
	}

	for _, tt := range tests {
		result = builtinStr(tt.input)
		if result.Inspect() != tt.expected {
			t.Errorf("wrong parsed result. expected=%s, got=%s", tt.expected, result.Inspect())
		}
	}
}
