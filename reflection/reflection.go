package reflection

import "reflect"

/* First: Easiest syntax for understanding:
func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}
*/

/* Second:
func walk(x interface{}, fn func(string)) {
	val := getValue(x)

	numberOfValues := 0
	// function variable, assigned to different fucntions based on the kind of value being processed
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())

	case reflect.Struct:
		numberOfValues = val.NumField()
		//val.Field is a method provided by the reflect package
		getField = val.Field // creates a function that retireves the field value at a guven index of the struct

	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		// val.Index is a method provded by the reflect package
		getField = val.Index // creates a function that retrieves the element value at a given index of the slice
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < numberOfValues; i++ {
		// Inteface method is called on the reflect.Value to get the actual value as an interface{}
		// walk is recursivley called with this value and the fn function
		walk(getField(i).Interface(), fn)
	}
}
*/

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

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
			// val.Recv() performs a receive operation on the channel represeented by val
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		// empty list of arguments to the function. This indicates that fucntion should be invoked with no arguments
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
