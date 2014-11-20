package objecttemplate

import (
	"reflect"

	"fknsrs.biz/p/dotty"
)

func Render(template interface{}, data interface{}) (res interface{}, err error) {
	_type := reflect.TypeOf(template)

	if _type == nil {
		res = template
	} else if _type.Kind() == reflect.String {
		tpl := reflect.ValueOf(template).String()

		if tpl[0:1] == "$" {
			res, err = dotty.Get(data, tpl[1:])
		} else {
			res = tpl
		}
	} else if _type.Kind() == reflect.Map && _type.Key().Kind() == reflect.String {
		tpl := reflect.ValueOf(template)
		obj := map[string]interface{}{}

		for _, k := range tpl.MapKeys() {
			if val, _err := Render(tpl.MapIndex(k).Interface(), data); _err != nil {
				err = _err
				return
			} else {
				obj[k.String()] = val
			}
		}

		res = obj
	} else if _type.Kind() == reflect.Array || _type.Kind() == reflect.Slice {
		tpl := reflect.ValueOf(template)
		arr := []interface{}{}

		for i := 0; i < tpl.Len(); i++ {
			if val, _err := Render(tpl.Index(i).Interface(), data); _err != nil {
				err = _err
				return
			} else {
				arr = append(arr, val)
			}
		}

		res = arr
	} else {
		res = template
	}

	return
}
