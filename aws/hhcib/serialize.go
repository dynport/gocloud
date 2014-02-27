package main

import (
	"net/url"
	"reflect"
	"strconv"
	"time"
)

func addValue(values url.Values, key string, value reflect.Value) {
	switch v := value.Interface().(type) {
	case string:
		if v != "" {
			values.Add(key, v)
		}
	case bool:
		if v {
			values.Add(key, "true")
		}
	case time.Time:
		if !v.IsZero() {
			values.Add(key, v.UTC().Format(time.RFC3339))
		}
	default:
		if value.Kind() == reflect.Ptr {
			urlValues(values, value.Elem().Interface(), key)
		}
	}
}

func urlValues(values url.Values, in interface{}, prefix string) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		sf := t.Field(i)
		name := sf.Tag.Get("aws")
		if prefix != "" {
			name = prefix + "." + name
		}
		if field.Kind() == reflect.Slice {
			for j := 0; j < field.Len(); j++ {
				addValue(values, name+"."+strconv.Itoa(j+1), field.Index(j))
			}
		} else {
			addValue(values, name, field)
		}
	}
}
