package check

import (
	"reflect"
	"strings"
)

//IsEmptyOrWhiteSpace checks string is "" or whitespace
func IsEmptyOrWhiteSpace(str string) bool {
	if IsEmpty(str) || len(strings.TrimSpace(str)) == 0 {
		return true
	}

	return false
}

//IsEmpty checks object is zero or nil
func IsEmpty(obj interface{}) bool {

	// get nil case out of the way
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
