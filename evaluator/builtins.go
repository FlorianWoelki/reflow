package evaluator

import (
	"github.com/florianwoelki/reflow/object/builtin"
)

var builtins = map[string]*builtin.Builtin{
	"len":    builtin.GetBuiltinByName("len"),
	"print":  builtin.GetBuiltinByName("print"),
	"str":    builtin.GetBuiltinByName("str"),
	"first":  builtin.GetBuiltinByName("first"),
	"last":   builtin.GetBuiltinByName("last"),
	"rest":   builtin.GetBuiltinByName("rest"),
	"pop":    builtin.GetBuiltinByName("pop"),
	"push":   builtin.GetBuiltinByName("push"),
	"map":    builtin.GetBuiltinByName("map"),
	"find":   builtin.GetBuiltinByName("find"),
	"filter": builtin.GetBuiltinByName("filter"),
	"delete": builtin.GetBuiltinByName("delete"),
}
