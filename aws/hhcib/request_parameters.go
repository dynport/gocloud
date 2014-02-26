package main

import "strings"

type RequestParameters []*Property

func (r RequestParameters) fields() []*TypeField {
	multiFields := map[string]map[string]string{}
	fields := []*TypeField{}
	for _, p := range r {
		name := p.Name
		field := &TypeField{Name: normalizeName(p.Name), Type: convertType(p.Type), Comments: Comments{"aws": strings.Replace(name, ".n", "", -1)}}
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
				theType := convertType(strings.Split(p.Type, ".")[0])
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
	for name, mapping := range multiFields {
		normalizedName := strings.TrimSpace(strings.Replace(name, ".", "", -1))
		t := &Type{Name: normalizedName}
		for subname, subType := range mapping {
			t.Fields = append(t.Fields, &TypeField{Name: strings.TrimSpace(strings.Replace(subname, ".", "", -1)), Type: subType, Comments: Comments{"aws": subname}})
		}
		fields = append(fields, &TypeField{CustomType: t, Name: pluralize(normalizeName(normalizedName)), Type: "[]*" + normalizedName, Comments: Comments{"aws": name}})
	}
	return fields
}
