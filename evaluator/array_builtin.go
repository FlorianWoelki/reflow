package evaluator

import "github.com/florianwoelki/reflow/object"

func First(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `first` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	if len(arr.Elements) > 0 {
		return arr.Elements[0]
	}

	return NULL
}

func Last(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `last` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		return arr.Elements[length-1]
	}

	return NULL
}

func Rest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `rest` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)
	if length > 0 {
		newElements := make([]object.Object, length-1)
		copy(newElements, arr.Elements[1:length])
		return &object.Array{Elements: newElements}
	}

	return NULL
}

func Pop(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `pop` must be ARRAY, got %s", args[0].Type())
	}

	array := args[0].(*object.Array)
	length := len(array.Elements)

	newElements := make([]object.Object, length-1)
	for i := 0; i < length-1; i++ {
		newElements[i] = array.Elements[i]
	}

	return &object.Array{Elements: newElements}
}

func Push(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, expected=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("argument to `push` must be ARRAY, got %s", args[0].Type())
	}

	arr := args[0].(*object.Array)
	length := len(arr.Elements)

	newElements := make([]object.Object, length+1)
	copy(newElements, arr.Elements)
	newElements[length] = args[1]

	return &object.Array{Elements: newElements}
}

func Map(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, expected=>=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("first argument to `map` must be ARRAY. got=%s", args[0].Type())
	}
	if args[1].Type() != object.FUNCTION_OBJ {
		return newError("second argument to `map` must be FUNCTION. got=%s", args[0].Type())
	}

	newElements := make([]object.Object, len(args)-2)
	copy(newElements, args[2:])

	return &object.Array{Elements: newElements}
}

func Find(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, expected=>=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("first argument to `find` must be ARRAY. got=%s", args[0].Type())
	}
	if args[1].Type() != object.FUNCTION_OBJ {
		return newError("second argument to `find` must be FUNCTION. got=%s", args[0].Type())
	}

	newElements := make([]object.Object, len(args)-2)
	copy(newElements, args[2:])

	foundIndex := -1
	for i, element := range newElements {
		if element, ok := element.(*object.Boolean); ok {
			if element.Value {
				foundIndex = i
				break
			}
		}
	}

	return &object.Integer{Value: int64(foundIndex)}
}

func Filter(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, expected=>=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return newError("first argument to `filter` must be ARRAY. got=%s", args[0].Type())
	}
	if args[1].Type() != object.FUNCTION_OBJ {
		return newError("second argument to `filter` must be FUNCTION. got=%s", args[0].Type())
	}

	array := args[0].(*object.Array)
	var filteredElements []object.Object
	for i := 2; i < len(args); i++ {
		element := args[i]
		if element, ok := element.(*object.Boolean); ok {
			if element.Value {
				filteredElements = append(filteredElements, array.Elements[i-2])
			}
		}
	}

	return &object.Array{Elements: filteredElements}
}
