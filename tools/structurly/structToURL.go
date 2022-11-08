package structurly

import (
	"fmt"
	"net/url"
	"reflect"
)

func ToURL(s interface{}) string {
	r := url.Values{}
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	for i := 0; i < t.NumField(); i++ {
		k := t.Field(i).Tag.Get("strly")
		if k != "" {
			field := v.Field(i)
			switch field.Kind() {
				case reflect.Ptr:
					if !field.IsNil() {
						r.Add(k, fmt.Sprintf("%v", field.Elem().Interface()))
					}
				default:
					t := fmt.Sprintf("%v", field.Interface())
					if t == "0" || t == "" { continue }
					r.Add(k, t)
			}
		}
	}

	return r.Encode()
}