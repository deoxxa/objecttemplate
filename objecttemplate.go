package objecttemplate

import (
	"fknsrs.biz/dotty"
)

func Render(template interface{}, data interface{}) (res interface{}, err error) {
	switch template.(type) {
	case []interface{}:
		tpl := template.([]interface{})
		arr := make([]interface{}, len(template.([]interface{})))
		res = arr

		for i := range tpl {
			if val, verr := Render(tpl[i], data); verr != nil {
				err = verr

				return
			} else {
				arr[i] = val
			}
		}
	case map[string]interface{}:
		tpl := template.(map[string]interface{})
		obj := map[string]interface{}{}
		res = obj

		for k := range tpl {
			switch tpl[k].(type) {
			case map[string]interface{}:
				if val, verr := Render(tpl[k], data); err != nil {
					err = verr

					return
				} else {
					obj[k] = val
				}
			case string:
				if tpl[k].(string)[0:1] == "$" {
					if val, verr := dotty.Get(data, tpl[k].(string)[1:]); verr != nil {
						err = verr

						return
					} else {
						obj[k] = val
					}
				} else {
					obj[k] = tpl[k]
				}
			default:
				obj[k] = tpl[k]
			}
		}
	default:
		res = template
	}

	return
}
