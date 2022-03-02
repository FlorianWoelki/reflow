package builtin

import "github.com/florianwoelki/reflow/object"

type BuiltinFunction func(args ...object.Object) object.Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() object.ObjectType {
	return object.BUILTIN_OBJ
}

func (b *Builtin) Inspect() string {
	return "builtin function"
}
