package builder

import "reflect"

type placeHolder bool

const PlaceHolder = placeHolder(true)

type Resource struct {
	TypeName string
	Menu     []MenuItem
	Output   map[string]interface{}
}

type MenuItem struct {
	Question string
	Options  []Option // If this is empty, it takes raw user input
	Output   map[string]interface{}
}

type Option struct {
	Name   string
	Output map[string]interface{}
}

func cp(in interface{}) interface{} {
	v := reflect.ValueOf(in)
	t := v.Type()

	switch t.Kind() {
	case reflect.Map:
		out := reflect.MakeMap(t)
		for _, key := range v.MapKeys() {
			val := reflect.ValueOf(cp(v.MapIndex(key).Interface()))

			out.SetMapIndex(key, val)
		}

		return out.Interface()
	case reflect.Array, reflect.Slice:
		out := reflect.MakeSlice(t, 0, v.Len())
		for i := 0; i < v.Len(); i++ {
			val := reflect.ValueOf(cp(v.Index(i).Interface()))

			out = reflect.Append(out, val)
		}

		return out.Interface()
	default:
		return in
	}
}

func munge(dst map[string]interface{}, src map[string]interface{}) {
	for key, value := range src {
		switch v := value.(type) {
		case map[string]interface{}:
			if _, ok := dst[key]; !ok || dst[key] == nil {
				dst[key] = make(map[string]interface{})
			}

			munge(dst[key].(map[string]interface{}), v)
		default:
			dst[key] = value
		}
	}
}

func (r Resource) Build() map[string]interface{} {
	output := cp(r.Output).(map[string]interface{})

	for _, item := range r.Menu {
		part := item.build()

		munge(output, part)
	}

	return map[string]interface{}{
		"Type":       r.TypeName,
		"Properties": output,
	}
}

func populate(output, set interface{}) {
	switch v := output.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if _, ok := value.(placeHolder); ok {
				v[key] = set
			} else {
				populate(value, set)
			}
		}
	case []interface{}:
		for i, value := range v {
			if _, ok := value.(placeHolder); ok {
				v[i] = set
			} else {
				populate(value, set)
			}
		}
	}
}

func (i MenuItem) build() map[string]interface{} {
	out := cp(i.Output)

	if len(i.Options) == 0 {
		populate(out, "NOT IMPLEMENTED")
	} else {
		part := cp(i.Options[0].Output)
		munge(out.(map[string]interface{}), part.(map[string]interface{}))
	}

	return out.(map[string]interface{})
}
