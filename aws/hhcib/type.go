package main

import "strings"

type Type struct {
	Name   string
	Fields []*TypeField
}

func (t *Type) String() string {
	lines := []string{"type " + t.Type() + " struct {"}
	for _, f := range t.Fields {
		lines = append(lines, "\t"+f.String())
	}
	lines = append(lines, "}\n")
	return strings.Join(lines, "\n")
}

func (t *Type) Type() string {
	return normalizeCustomType(t.Name)
}

const pointer = "*"

func convertType(s string) string {
	s = strings.TrimSpace(s)
	switch strings.ToLower(s) {
	case "xsd:string":
		return "string"
	case "integer", "long", "integer (16-bit unsigned)":
		return "int"
	case "double":
		return "float64"
	case "boolean":
		return "bool"
	case "empty element":
		return "struct{}"
	case "datetime":
		return "time.Time"
	case "[]string", "struct{}", "int", "string", "bool":
		return strings.ToLower(s)
	case "":
		return "struct{}"
	default:
		return pointer + strings.TrimSuffix(s, "Type")
	}
}
