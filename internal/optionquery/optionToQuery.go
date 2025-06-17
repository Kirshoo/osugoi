package optionquery

import (
	"net/url"
	"reflect"
	"strings"
)

func Convert[T any](options T) url.Values {
	v := reflect.ValueOf(options)
	t := reflect.TypeOf(options)

	values := url.Values{}

	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		tag := fieldType.Tag.Get("query")
		if tag == "" || tag == "-" {
			tag = strings.ToLower(fieldType.Name)
		}

		values.Add(tag, fieldValue.String())
	}

	return values
}
