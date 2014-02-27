package main

import "strings"

type RequestParameters []*Property

var requestTypeMapping = map[string]string{
	"NetworkInterface": "RequestNetworkInterface",
}

func mapRequestType(t string) string {
	if mapped, ok := requestTypeMapping[t]; ok {
		return mapped
	}
	return t
}

func (r RequestParameters) fields() []*TypeField {
	multiFields := MultiFieldMap{}
	fields := []*TypeField{}
	for _, p := range r {
		name := p.Name
		t := mapRequestType(convertType(p.Type))
		field := &TypeField{Name: normalizeName(p.Name), Type: t, Comments: Comments{"aws": strings.Replace(name, ".n", "", -1)}}
		if strings.HasSuffix(name, ".n") {
			field.Name = pluralize(field.Name)
			field.Type = "[]" + field.Type
		} else if strings.Contains(name, ".n.") {
			parts := strings.Split(name, ".n.")
			if len(parts) > 1 {
				name := parts[0]
				subName := parts[1]
				if multiFields[name] == nil {
					multiFields[name] = map[string]string{}
				}
				theType := convertType(strings.Split(t, ".")[0])
				if strings.HasSuffix(subName, ".m") {
					subName = strings.TrimSuffix(subName, ".m") + "s"
					theType = "[]" + theType
				}
				multiFields[name][subName] = theType
			}
			continue
		}
		fields = append(fields, field)
	}
	return append(fields, multiFields.fields()...)
}
