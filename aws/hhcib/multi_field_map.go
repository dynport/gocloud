package main

import "strings"

type MultiFieldMap map[string]map[string]string

func (mf MultiFieldMap) fields() []*TypeField {
	fields := []*TypeField{}
	for name, mapping := range mf {
		normalizedName := strings.TrimSpace(strings.Replace(name, ".", "", -1))
		p := RequestParameters{}
		for subname, subType := range mapping {
			p = append(p, &Property{Name: strings.Replace(subname, ".m.", ".n.", -1), Type: mapRequestType(subType)})
		}
		typ := mapRequestType(normalizedName)
		t := &Type{Name: mapRequestType(normalizedName), Fields: p.fields()}
		fields = append(fields, &TypeField{CustomType: t, Name: pluralize(normalizeName(normalizedName)), Type: "[]*" + typ, Comments: Comments{"aws": name}})
	}
	return fields
}
