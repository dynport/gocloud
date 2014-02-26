package main

import "log"

type Operation struct {
	Name   string `yaml:":name"`
	Method string `yaml:":method"`
	Inputs Input  `yaml:":inputs"`
}

type TypeMap map[interface{}]interface{}

func (t TypeMap) YamlType() *YamlType {
	if tmp, ok := t[":list"]; ok {
		switch tmp := tmp.(type) {
		case []interface{}:
			if len(tmp) > 0 {
				switch tmp := tmp[0].(type) {
				case string:
					return &YamlType{List: true, Type: tmp[1:]}
				case map[interface{}]interface{}:
					if s, ok := tmp[":structure"]; ok {
						return &YamlType{List: true, Struct: true, StructFields: castStruct(s)}
					}
				}
			}
		}
	}
	return nil
}

func castStruct(i interface{}) []*YamlType {
	s, ok := i.(map[interface{}]interface{})
	if !ok {
		return nil
	}
	types := []*YamlType{}
	for k, v := range s {
		values, ok := v.([]interface{})
		if !ok {
			continue
		}
		t := extractYamlType(values)
		if t != nil {
			if s, ok := k.(string); ok {
				t.Name = s
				types = append(types, t)
			}
		}
	}
	return types
}

type Structure map[string][]interface{}

type YamlType struct {
	Name         string
	Type         string
	List         bool
	Struct       bool
	StructFields []*YamlType
}

func extractYamlType(in []interface{}) *YamlType {
	if len(in) == 0 {
		return nil
	}
	switch t := in[0].(type) {
	case string:
		return &YamlType{Type: t[1:]}
	case map[interface{}]interface{}:
		return TypeMap(t).YamlType()
	default:
		log.Printf("ERROR: %T %#v", t, t)
		return nil
	}
}

func (op *Operation) YamlTypes() []*YamlType {
	fields := []*YamlType{}
	for name, input := range op.Inputs {
		t := extractYamlType(input)
		if t != nil {
			t.Name = name
			fields = append(fields, t)
		}
	}
	return fields
}

type Input map[string][]interface{}

func (input Input) Fields() []*TypeField {
	fields := []*TypeField{}
	for name, _ := range input {
		fields = append(fields, &TypeField{Name: name})
	}
	return fields
}

type Outputs struct {
	Children map[string]*Outputs `yaml:":children,omitempty"`
	Type     string              `yaml:":type,omitempty"`
	Ignore   bool                `yaml:":ignore"`
	List     bool                `yaml:":list"`
}
