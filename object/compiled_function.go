package object

import (
	"fmt"

	"github.com/florianwoelki/reflow/code"
)

type CompiledFunction struct {
	Instructions code.Instructions
}

func (cf *CompiledFunction) Type() ObjectType {
	return COMPILED_FUNCTION_OBJ
}

func (cf *CompiledFunction) Inspect() string {
	return fmt.Sprintf("CompiledFunction[%p]", cf)
}
