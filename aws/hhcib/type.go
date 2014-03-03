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
		return "*IntValue"
	case "double":
		return "*FloatValue"
	case "boolean":
		return "*BoolValue"
	case "empty element":
		return "struct{}"
	case "datetime":
		return "time.Time"
	case "[]string", "struct{}", "*IntValue", "string", "*BoolValue", "*FloatValue":
		return strings.ToLower(s)
	case "":
		return "struct{}"
	default:
		s := strings.TrimSuffix(s, "Type")
		if strings.HasPrefix(s, pointer) {
			return s
		}
		return pointer + s
	}
}
