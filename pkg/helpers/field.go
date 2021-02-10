package helpers

import "reflect"

func DeepFields(iface interface{}) []reflect.StructField {
	fields := make([]reflect.StructField, 0)
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)

	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)

		switch v.Kind() {
		case reflect.Struct:
			fields = append(fields, DeepFields(v.Interface())...)
		default:
			fields = append(fields, ifv.Type().Field(i))
		}
	}

	return fields
}
