package walking

import (
	"reflect"
)

// takes in a blob or interface and a function which serves well for testing
// we can insert our own function in tests to ensure the reflection logic is working
func walk(x interface{}, fn func(string)) {
	// wraps x in reflection
	val := getValue(x)

	// removes the reflection wrapper and calls walk recursively
	// this is required for structs and other data types that can contain more than one value
	// we might need to iterate on each value and run the function on individual values
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	// uses reflection to determine the value "type" or Kind
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}

// helper to wrap values in reflection
// also has specific handling for pointers to return the pointed value
func getValue(blob interface{}) reflect.Value {
	val := reflect.ValueOf(blob)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
