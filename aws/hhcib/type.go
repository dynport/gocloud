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

func convertType(s string) string {
	switch strings.ToLower(s) {
	case "string", "xsd:string":
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
	case "":
		return "struct{}"
	default:
		return "*" + strings.TrimSuffix(s, "Type")
	}
}

//func (t *Type) Name() string {
//	return t.name
//}
//
//func (t *Type) Parse(doc xml.Node) error {
//	a, e := parseDocuNode(doc)
//	if e != nil {
//		return e
//	}
//	t.name = strings.TrimSuffix(a.Type, "Type")
//	contents := a.AllProperties["Contents"]
//	for _, field := range contents {
//		t.fields = append(t.fields, &Field{Name: field.Name, Type: field.Type, XmlName: field.Name})
//	}
//	return nil
//}
//
//func (a *Type) String() string {
//	out := []string{}
//	out = append(out, "type "+a.name+" struct {")
//	for _, p := range a.fields {
//		t := p.Type
//		if t == "" {
//			t = "struct{}"
//		}
//		t = convertType(t)
//		comments := []string{}
//		if p.AttributeName != "" {
//			comments = append(comments, fmt.Sprintf(`aws:"%s"`, strings.TrimSpace(p.AttributeName)))
//		}
//		if p.XmlName != "" {
//			comments = append(comments, fmt.Sprintf(`xml:"%s"`, strings.TrimSpace(p.XmlName)))
//		}
//		line := "\t" + strings.Title(strings.Replace(strings.TrimSpace(p.Name), ".", "", -1)) + " " + t
//		if len(comments) > 0 {
//			line += " `" + strings.Join(comments, " ") + "`"
//		}
//		if p.CustomField {
//			line += " // custom"
//		}
//		out = append(out, line) // " "+fmt.Sprintf("`"+`xml:"%s,omitempty"`+"`", p.Name))
//	}
//	out = append(out, "}")
//	return strings.Join(out, "\n") + "\n"
//}
