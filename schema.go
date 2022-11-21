package golang

import (
	"strings"

	"github.com/schema-cafe/go-types/form"
)

func Schema(name string, fields []form.Field) string {
	s := strings.Builder{}
	s.WriteString("type ")
	s.WriteString(name)
	s.WriteString(" struct {\n")
	for _, f := range fields {
		s.WriteString("\t")
		s.WriteString(f.Name)
		s.WriteString(" ")
		s.WriteString(Type(f.Type))
		s.WriteString("\n")
	}
	s.WriteString("}\n")
	return s.String()
}
