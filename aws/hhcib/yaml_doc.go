package main

import "launchpad.net/goyaml"

type YamlDoc struct {
	ApiVersion string       `yaml:":api_version"`
	Operations []*Operation `yaml:":operations"`
}

func (doc *YamlDoc) Load(b []byte) error {
	return goyaml.Unmarshal(b, doc)
}
