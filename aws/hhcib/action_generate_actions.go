package main

type GenerateActions struct {
}

func (g *GenerateActions) Run() error {
	doc, e := loadDoc(ec2ApiRoot)
	if e != nil {
		return e
	}
	links, e := extractLinks(doc)
	if e != nil {
		return e
	}
	logger.Printf("generating %d actions", len(links))
	types := []*Type{}
	customTypes := map[string]*Type{}
	for _, link := range links {
		logger.Printf("writing %s", link.Name)
		doc, e := loadDoc(link.Url)
		if e != nil {
			return e
		}
		a, e := parseDocuNode(doc)
		if e != nil {
			return e
		}
		ac := &Action{
			Name:              a.Type,
			RequestParameters: a.RequestParameters(),
		}
		t := ac.RequestType()
		for _, f := range t.Fields {
			if f.CustomType != nil {
				customTypes[f.CustomType.Name] = f.CustomType
			}
		}
		types = append(types, t)
	}
	if len(customTypes) > 0 {
		logger.Print("CUSTOM TYPES")
		for name, t := range customTypes {
			logger.Printf(" - %s", name)
			types = append(types, t)
		}
	}
	return writeTypes("generated", "generated/ec2_actions.go", types)
}

func init() {
	router.Register("ec2/generate/actions", &GenerateActions{}, "Generate Actions")
}
