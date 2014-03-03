package main

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type IntValue struct {
	Value int64
}

type FloatValue struct {
	Value float64
}

type BoolValue struct {
	Value bool
}

func addValue(values url.Values, key string, value reflect.Value) {
	switch v := value.Interface().(type) {
	case string:
		if v != "" {
			values.Add(key, v)
		}
	case *IntValue:
		if v != nil {
			values.Add(key, strconv.FormatInt(v.Value, 10))
		}
	case *FloatValue:
		if v != nil {
			values.Add(key, fmt.Sprintf("%.6f", v.Value))
		}
	case *BoolValue:
		if v != nil {
			values.Add(key, fmt.Sprintf("%t", v.Value))
		}
	case time.Time:
		if !v.IsZero() {
			values.Add(key, v.UTC().Format(time.RFC3339))
		}
	default:
		if value.Kind() == reflect.Ptr && !value.IsNil() {
			urlValues(values, value.Elem().Interface(), key)
		}
	}
}

func urlValues(values url.Values, in interface{}, prefix string) {
	if prefix == "" {
		typ := fmt.Sprintf("%T", in)
		parts := strings.Split(typ, ".")
		last := parts[len(parts)-1]
		values.Set("Action", last)
		values.Set("Version", EC2_VERSION)
	}
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	switch v.Kind() {
	case reflect.Struct:
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
	case reflect.Int:

	default:
		// unknwon
	}
}
