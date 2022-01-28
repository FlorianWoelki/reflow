package evaluator

import (
	"testing"

	"github.com/florianwoelki/reflow/object"
)

func TestBuiltinFilter(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "filter(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "filter()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "filter([1, 2, 3])",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "filter([1, 2, 3, 4], 5)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "filter([1, 2, 3, 4], fn(x) { x < 3 })",
			expected:        "[1, 2]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "filter([1, 2, 3, 4], fn(x) { true })",
			expected:        "[1, 2, 3, 4]",
			expectedObjType: object.ARRAY_OBJ,
		},
	}

	testInput(t, "filter", tests)
}

func TestBuiltinFind(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "find(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "find()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "find([1, 2, 3])",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "find([1, 2, 3, 4], 5)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "find([1, 2, 3, 4], fn(x) { x == 2 })",
			expected:        "1",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "find([1, 2, 3, 4], fn(x) { true })",
			expected:        "0",
			expectedObjType: object.INTEGER_OBJ,
		},
	}

	testInput(t, "find", tests)
}

func TestBuiltinPush(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "push(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "push()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "push([1, 2, 3])",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "push([1, 2, 3, 4], 5)",
			expected:        "[1, 2, 3, 4, 5]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "push([1], 2)",
			expected:        "[1, 2]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "push([], 1)",
			expected:        "[1]",
			expectedObjType: object.ARRAY_OBJ,
		},
	}

	testInput(t, "push", tests)
}

func TestBuiltinPop(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "pop(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "pop()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "pop([1, 2, 3, 4], 4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "pop([1, 2, 3, 4])",
			expected:        "[1, 2, 3]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "pop([1])",
			expected:        "[]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "pop([])",
			expected:        "[]",
			expectedObjType: object.ARRAY_OBJ,
		},
	}

	testInput(t, "pop", tests)
}

func TestBuiltinRest(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "rest(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "rest()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "rest([1, 2, 3, 4], 4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "rest([1, 2, 3, 4])",
			expected:        "[2, 3, 4]",
			expectedObjType: object.ARRAY_OBJ,
		},
		{
			input:           "last([1])",
			expected:        "1",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "last([])",
			expected:        "null",
			expectedObjType: object.NULL_OBJ,
		},
	}

	testInput(t, "rest", tests)
}

func TestBuiltinLast(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "last(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "last()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "last([1, 2, 3, 4], 4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "last([1, 2, 3, 4])",
			expected:        "4",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "last([1])",
			expected:        "1",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "last([])",
			expected:        "null",
			expectedObjType: object.NULL_OBJ,
		},
	}

	testInput(t, "last", tests)
}

func TestBuiltinFirst(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "first(4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "first()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "first([1, 2, 3, 4], 4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "first([1, 2, 3, 4])",
			expected:        "1",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "first([1])",
			expected:        "1",
			expectedObjType: object.INTEGER_OBJ,
		},
		{
			input:           "first([])",
			expected:        "null",
			expectedObjType: object.NULL_OBJ,
		},
	}

	testInput(t, "first", tests)
}

func TestBuiltinMap(t *testing.T) {
	tests := []builtinTest{
		{
			input:           "map(4, fn(x) { x * 2})",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "map([1, 2, 3, 4])",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "map()",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "map([1, 2, 3, 4], 4)",
			expected:        "",
			expectedObjType: object.ERROR_OBJ,
		},
		{
			input:           "map([1, 2, 3], fn(x) { x * 2})",
			expected:        "[2, 4, 6]",
			expectedObjType: object.ARRAY_OBJ,
		},
	}

	testInput(t, "map", tests)
}

func TestBuiltinsLength(t *testing.T) {
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

type builtinTest struct {
	input           string
	expected        string
	expectedObjType object.ObjectType
}

func testInput(t *testing.T, builtinFn string, tests []builtinTest) {
	for _, tt := range tests {
		output := testEval(tt.input)
		if output.Type() != tt.expectedObjType {
			t.Errorf("builtin %s function test returned not the expected object type. expected=%s, got=%s", builtinFn, tt.expectedObjType, output.Type())
		}

		if output.Type() == object.ERROR_OBJ {
			continue
		}

		if output.Inspect() != tt.expected {
			t.Errorf("builtin %s function test returned not the correct value. expected=%s, got=%s", builtinFn, tt.expected, output.Inspect())
		}
	}
}
