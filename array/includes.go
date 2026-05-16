package array

import (
	"reflect"
)

func Includes[S ~[]T, T any](slice S, element T, start *int) bool {

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

// The Includes() method determines whether the array includes a certain element.
func (a Array[T]) Includes(element T, start *int) bool {
	return Includes(a, element, start)
}
