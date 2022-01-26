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
	"first":  {Fn: First},
	"last":   {Fn: Last},
	"rest":   {Fn: Rest},
	"pop":    {Fn: Pop},
	"push":   {Fn: Push},
	"map":    {Fn: Map},
	"find":   {Fn: Find},
	"filter": {Fn: Filter},
}
