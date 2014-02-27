package main

import (
	"fmt"
	"strings"
)

func newTypeField(name, fieldType string, comments Comments) *TypeField {
	return &TypeField{Name: name, Type: fieldType, Comments: comments}
}

type TypeField struct {
	Name       string
	Type       string
	Comments   Comments
	CustomType *Type
}

func (tf *TypeField) String() string {
	out := strings.TrimSpace(tf.Name) + " " + tf.Type
	comments := []string{}
	for key, value := range tf.Comments {
		comments = append(comments, fmt.Sprintf("%s:%q", key, value))
	}
	if len(comments) > 0 {
		out += " `" + strings.Join(comments, " ") + "`"
	}
	return out
}
