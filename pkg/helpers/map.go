package helpers

import (
	"github.com/fatih/structs"
	"reflect"
	"strings"
)

func Lower(f interface{}) interface{} {
	switch f := f.(type) {
	case []interface{}:
		for i := range f {
			f[i] = Lower(f[i])
		}
		return f
	case map[string]interface{}:
		lf := make(map[string]interface{}, len(f))
		for k, v := range f {
			lf[strings.ToLower(k)] = Lower(v)
		}
		return lf
	default:
		return f
	}
}

func ConvertToObject(obj interface{}, attributes map[string]string) interface{} {
	val := reflect.ValueOf(obj)
	names := structs.Names(obj)
	for _, name := range names {
		lowerName := strings.ToLower(name)
		if attributeValue, ok := attributes[lowerName]; ok {
			fl := val.FieldByName(name)
			fl.SetString(attributeValue)
		}
	}
	return obj
}
