package main

type DocumentationPage struct {
	Name             string
	DocumentationUrl string
	Type             string
	AllProperties    map[string][]*Property
}

func (a *DocumentationPage) RequestParameters() RequestParameters {
	return RequestParameters(a.AllProperties["Request Parameters"])
}

func (a *DocumentationPage) ReturnValues() []*Property {
	return a.AllProperties["ReturnValues"]
}

func (a *DocumentationPage) ResponseElements() []*Property {
	return a.AllProperties["ResponseElements"]
}

func (a *DocumentationPage) Properties() []*Property {
	return a.AllProperties["Properties"]
}

func (a *DocumentationPage) Contents() ([]*Property, error) {
	return nil, nil
}
