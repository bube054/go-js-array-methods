package array

import (
	"reflect"
)

// The IndexOf() function returns the first index (position) of a specified value and a possible error due to indexes out of range. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right).
// NOTE: ❌ POTENTIAL PERFORMANCE TRAP
// Forces the caller to allocate an integer on the heap just to pass an index.
// Consider using variadic parameters better performance
func LastIndexOf[S ~[]T, T comparable](slice S, element T, start *int) int {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return -1
	}

	startIndex := sliceLength - 1
	if start != nil {
		startIndex = convertIndexSimply(sliceLength, *start)

		if startIndex == -1 {
			return -1
		}
	}

	for i := startIndex; i >= 0; i-- {
		if slice[i] == element {
			return i
		}
	}

	return -1
}

// The LastIndexOf() method returns the last index at which a given element can be found in the array.
// Deprecated: Consider using ArrayComp.LastIndexOf instead.
func (a Array[T]) LastIndexOf(element T, start *int) int {
	sliceLength := len(a)

	startIndex := sliceLength - 1

	if start != nil {
		startIndex = convertIndexSimply(sliceLength, *start)
		if startIndex == -1 {
			return -1
		}
	}

	for i := startIndex; i >= 0; i-- {
		value := a[i]

		if reflect.DeepEqual(element, value) {
			return i
		}
	}

	return -1
}

// The LastIndexOf() method returns the last index at which a given element can be found in the array.
func (a ArrayComp[T]) LastIndexOf(element T, start ...int) int {
	index := OptionalParam(start, 0)
	return LastIndexOf(a, element, &index)
}
