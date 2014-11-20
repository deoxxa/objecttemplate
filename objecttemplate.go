package objecttemplate

import (
	"reflect"

	"fknsrs.biz/p/dotty"
)

func Render(template interface{}, data interface{}) (interface{}, error) {
	t := reflect.TypeOf(template)

	if t == nil {
		return template, nil
	} else if t.Kind() == reflect.String {
		s := reflect.ValueOf(template).String()

		if s[0:1] == "$" {
			return dotty.Get(data, s[1:])
		} else {
			return s, nil
		}
	} else if t.Kind() == reflect.Map && t.Key().Kind() == reflect.String {
		m := reflect.ValueOf(template)
		r := map[string]interface{}{}

		for _, k := range m.MapKeys() {
			if v, err := Render(m.MapIndex(k).Interface(), data); err != nil {
				return nil, err
			} else {
				r[k.String()] = v
			}
		}

		return r, nil
	} else if t.Kind() == reflect.Array || t.Kind() == reflect.Slice {
		a := reflect.ValueOf(template)
		r := []interface{}{}

		for i := 0; i < a.Len(); i++ {
			if v, err := Render(a.Index(i).Interface(), data); err != nil {
				return nil, err
			} else {
				r = append(r, v)
			}
		}

		return r, nil
	}

	return template, nil
}
