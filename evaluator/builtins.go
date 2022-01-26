package evaluator

import (
	"fmt"

	"github.com/florianwoelki/reflow/object"
)

var builtins = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported. got=%s", args[0].Type())
			}
		},
	},
	"first": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"last": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"rest": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"push": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"map": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"str": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, expected=1", len(args))
			}
			if args[0].Type() == object.ARRAY_OBJ {
				return newError("argument to `str` cannot be ARRAY")
			}

			return &object.String{Value: args[0].Inspect()}
		},
	},
	"find": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"filter": {
		Fn: func(args ...object.Object) object.Object {
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
			newElements := make([]object.Object, len(args)-2)
			copy(newElements, args[2:])

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
		},
	},
}
