package golang

import "github.com/schema-cafe/go-types/ast"

func Type(t ast.Type) string {
	if t.IsArray {
		return "[]" + Identifier(t.BaseType)
	}
	if t.IsMap {
		return "map[string]" + Identifier(t.BaseType)
	}
	if t.IsPointer {
		return "*" + Identifier(t.BaseType)
	}
	return Identifier(t.BaseType)
}
