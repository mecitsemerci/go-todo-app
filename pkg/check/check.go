package check

import (
	"reflect"
	"strings"
)

//IsEmptyOrWhiteSpace checks string is "" or whitespace
func IsEmptyOrWhiteSpace(str string) bool {
	return IsEmpty(str) || len(strings.TrimSpace(str)) == 0 
}

//IsEmpty checks object is zero or nil
func IsEmpty(obj any) bool {
	if obj == nil {
		return true
	}

	objValue := reflect.ValueOf(obj)

	switch objValue.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	case reflect.Ptr:
		if objValue.IsNil() {
			return true
		}
		ref := objValue.Elem().Interface()
		return IsEmpty(ref)
	default:
		zero := reflect.Zero(objValue.Type())
		return reflect.DeepEqual(obj, zero.Interface())
	}
}
