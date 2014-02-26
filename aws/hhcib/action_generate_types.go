package main

import (
	"io"
	"os"
	"strings"
)

type GenerateTypes struct {
}

func (g *GenerateTypes) Run() error {
	doc, e := loadDoc(ec2TypesUrl)
	if e != nil {
		return e
	}
	links, e := extractLinks(doc)
	if e != nil {
		return e
	}
	logger.Printf("extracted %d links", len(links))

	types := []*Type{}
	for _, link := range links {
		doc, e := loadDoc(link.Url)
		if e != nil {
			return e
		}
		a, e := parseDocuNode(doc)
		if e != nil {
			return e
		}
		fields := []*TypeField{}

		for _, p := range a.AllProperties["Contents"] {
			fields = append(fields, p.toTypeDefinition())
		}
		types = append(types, &Type{Name: a.Type, Fields: fields})
	}
	return writeTypes("generated", "generated/types.go", types)
}

func writeTypes(pkgName string, path string, types []*Type) error {
	os.Mkdir("generated", 0755)
	f, e := os.Create(path)
	if e != nil {
		return e
	}
	defer f.Close()
	typeLines := []string{}
	for _, t := range types {
		typeLines = append(typeLines, t.String())
	}
	typesString := strings.Join(typeLines, "\n")
	all := []string{"package generated\n"}

	if strings.Contains(typesString, "time.Time") {
		all = append(all, `import "time"`+"\n")
	}
	all = append(all, typesString)
	_, e = io.WriteString(f, strings.Join(all, "\n"))
	return e
}

func init() {
	router.Register("ec2/generate/types", &GenerateTypes{}, "Generate EC2 Types")
}
