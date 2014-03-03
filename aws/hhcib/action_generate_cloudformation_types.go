package main

import (
	"fmt"
	"log"
	"strings"
)

type GenerateCloudformationTypes struct {
}

var cloudFormationResource = "http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html"

func (g *GenerateCloudformationTypes) Run() error {
	types := []*Type{}

	for _, r := range []string{cloudFormationResource} {
		logger.Printf("generating types")
		doc, e := loadDoc(r)
		if e != nil {
			return e
		}
		links, e := extractLinks(doc, urlRoot(r))
		if e != nil {
			return e
		}
		logger.Printf("extracted %d links", len(links))

		for _, l := range links {
			log.Print(l.Url)
			doc, e := loadDoc(l.Url)
			if e != nil {
				return fmt.Errorf("error loading doc for ulr=%s: error=%q", l.Url, e)
			}
			a, e := parseDocuNode(doc)
			if e != nil {
				return e
			}
			fields := []*TypeField{}
			for _, p := range a.AllProperties["Properties"] {
				fields = append(fields, &TypeField{
					Name:     strings.Title(p.Name),
					Type:     "interface{}",
					Comments: Comments{"json": p.Name + ",omitempty"},
				})
			}
			types = append(types, &Type{Name: normalizeCloudformationType(a.Type), Fields: fields})
		}
	}
	return writeTypes("main", "cf_types_generated.go", types)
}

func normalizeCloudformationType(t string) string {
	return strings.Replace(strings.Replace(strings.TrimPrefix(t, "AWS::"), "::", "", -1), " ", "", -1)
}

func init() {
	router.Register("cf/generate/types", &GenerateCloudformationTypes{}, "Generate Cloudformation Types")
}
