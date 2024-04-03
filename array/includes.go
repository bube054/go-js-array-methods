package array

import (
	"reflect"
)

func Includes[T comparable](slice []T, element T, start *int) bool {

	var (
		startIndex int
		err        error
	)

	if start == nil {
		startIndex, err = ConvertIndex(slice, 0, "start index")
	} else {
		startIndex, err = ConvertIndex(slice, *start, "start index")
	}

	if err != nil {
		return false
	}

	sliceLength := len(slice)

	for i := startIndex; i < sliceLength; i++ {
		value := slice[i]

		if reflect.DeepEqual(element, value) {
			return true
		}
	}

	return false
}
