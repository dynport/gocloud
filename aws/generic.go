package aws

import (
	"fmt"
	"net/url"
	"reflect"
)

// beging marker types
type Version struct {
}

type Action struct {
}

// end marker types

func ParamsForAction(i interface{}) (url.Values, error) {
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	values := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		awsTag := fieldType.Tag.Get("aws")
		if awsTag != "" {
			name := awsTag
			switch casted := v.Field(i).Interface().(type) {
			case int:
				if casted > 0 {
					values.Add(name, fmt.Sprintf("%d", casted))
				}
			case string:
				if casted != "" {
					values.Add(name, casted)
				}
			case Version:
				values.Add("Version", awsTag)
			case Action:
				values.Add("Action", awsTag)
			default:
				return nil, fmt.Errorf("unable to handle %s (%T)", casted, casted)
			}
		}
	}
	return values, nil
}
