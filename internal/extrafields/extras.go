package extrafields

import (
	"reflect"
)

type WithExtras struct {
	ExtraFields map[string]any `json:"-"`
}

func ExtractKnownFields(v interface{}) map[string]struct{} {
	known := make(map[string]struct{})
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		if tag == "-" || tag == "" {
			continue
		}

		name := tag
		if comma := indexComma(tag); comma != -1 {
			name = tag[:comma]
		}
		if name == "" {
			name = field.Name
		}

		known[name] = struct{}{}
	}

	return known
}

func indexComma(tag string) int {
	for i := 0; i < len(tag); i++ {
		if tag[i] == ',' {
			return i
		}
	}

	return -1
}
