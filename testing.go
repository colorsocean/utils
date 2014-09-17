package utils

import (
	"reflect"
)

const (
	success = ""
)

func ShouldBeDeepEqual(actual interface{}, expected ...interface{}) string {
	if reflect.DeepEqual(actual, expected[0]) {
		return success
	} else {
		return "Assertion: Objects are not deep equal, but they must!"
	}
}

func ShouldNotBeDeepEqual(actual interface{}, expected ...interface{}) string {
	if !reflect.DeepEqual(actual, expected[0]) {
		return success
	} else {
		return "Assertion: Objects must be not deep equal, but they is!"
	}
}
