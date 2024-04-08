package array

import (
	"reflect"
)

// The IndexOf() function returns the first index (position) of a specified value and a possible error due to indexes out of range. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right).
func LastIndexOf[T comparable](slice []T, element T, start *int) int {
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
		return -1
	}

	sliceLength := len(slice)

	for i := sliceLength - 1; i >= startIndex; i-- {
		value := slice[i]

		if reflect.DeepEqual(element, value) {
			return i
		}
	}

	return -1
}
