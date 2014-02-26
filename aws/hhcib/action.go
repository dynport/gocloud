package main

type Action struct {
	Name              string
	RequestParameters RequestParameters
	ResponseElements  []*Property
}

func (a *Action) RequestType() *Type {
	fields := []*TypeField{}
	for _, f := range a.RequestParameters.fields() {
		fields = append(fields, f)
	}
	return &Type{Name: a.Name, Fields: fields}
}
