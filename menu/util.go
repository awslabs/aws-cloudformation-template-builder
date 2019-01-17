package menu

import (
	"reflect"

	"github.com/iancoleman/orderedmap"
)

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

func mungeMap(dst map[string]interface{}, src map[string]interface{}) {
	for key, value := range src {
		switch v := value.(type) {
		case map[string]interface{}:
			if _, ok := dst[key]; !ok || dst[key] == nil {
				dst[key] = make(map[string]interface{})
			}

			mungeMap(dst[key].(map[string]interface{}), v)
		default:
			dst[key] = value
		}
	}
}

func mungeOrderedMap(dst, src *orderedmap.OrderedMap) {
	for _, key := range src.Keys() {
		value, _ := src.Get(key)
		switch v := value.(type) {
		case orderedmap.OrderedMap:
			dst_v, ok := dst.Get(key)
			if !ok {
				dst_v = *orderedmap.New()
			}

			dst_om := dst_v.(orderedmap.OrderedMap)

			mungeOrderedMap(&dst_om, &v)

			dst.Set(key, dst_om)
		default:
			dst.Set(key, value)
		}
	}
}
