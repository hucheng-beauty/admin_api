package unique

import "reflect"

func String(strSlice []string) []string {
	m := make(map[string]bool)
	var list []string
	for _, v := range strSlice {
		if _, ok := m[v]; !ok {
			m[v] = true
			list = append(list, v)
		}
	}
	return list
}

func Int(intSli []int) []int {
	m := make(map[int]bool)
	var list []int
	for _, v := range intSli {
		if _, ok := m[v]; !ok {
			m[v] = true
			list = append(list, v)
		}
	}
	return list
}

func Any(data interface{}) interface{} {
	in := reflect.ValueOf(data)
	if in.Kind() != reflect.Slice && in.Kind() != reflect.Array {
		return data
	}

	m := make(map[interface{}]bool)

	out := reflect.MakeSlice(in.Type(), 0, in.Len())
	for i := 0; i < in.Len(); i++ {
		value := in.Index(i)

		if _, ok := m[value.Interface()]; !ok {
			out = reflect.Append(out, in.Index(i))
			m[value.Interface()] = true
		}
	}
	return out.Interface()
}
