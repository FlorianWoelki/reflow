package builtin

import (
	"fmt"

	"github.com/florianwoelki/reflow/object"
)

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}

	return nil
}

func builtinLen(args ...object.Object) object.Object {
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
}

func builtinPrint(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return nil
}

func builtinStr(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, expected=1", len(args))
	}
	if args[0].Type() == object.ARRAY_OBJ {
		return newError("argument to `str` cannot be ARRAY")
	}

	return &object.String{Value: args[0].Inspect()}
}

func builtinDelete(args ...object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, expected=2", len(args))
	}
	if args[0].Type() != object.HASH_OBJ {
		return newError("first argument to `delete` should be a hash")
	}

	hash := args[0].(*object.Hash)
	k := args[1].(object.Hashable)
	_, ok := hash.Pairs[k.HashKey()]
	if !ok {
		return newError("hash key `%d` not found", k.HashKey().Value)
	}

	delete(hash.Pairs, k.HashKey())

	return nil
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{"len", &Builtin{Fn: builtinLen}},
	{"print", &Builtin{Fn: builtinPrint}},
	{"str", &Builtin{Fn: builtinStr}},
	{"first", &Builtin{Fn: builtinFirst}},
	{"last", &Builtin{Fn: builtinLast}},
	{"rest", &Builtin{Fn: builtinRest}},
	{"pop", &Builtin{Fn: builtinPop}},
	{"push", &Builtin{Fn: builtinPush}},
	/*{"map", &Builtin{Fn: builtinMap}},
	{"find", &Builtin{Fn: builtinFind}},
	{"filter", &Builtin{Fn: builtinFilter}},*/
	{"delete", &Builtin{Fn: builtinDelete}},
}
