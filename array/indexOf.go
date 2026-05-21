package array

import (
	"reflect"
)

// IndexOf returns the first index (position) of a specified value. The IndexOf() function returns -1 if the value is not found. The IndexOf() function starts at a specified index and searches from left to right (from the given start postion to the end of the array). By default the search starts at the first element and ends at the last. Negative start values counts from the last element (but still searches from left to right).
// NOTE: ❌ POTENTIAL PERFORMANCE TRAP
// Forces the caller to allocate an integer on the heap just to pass an index.
// Consider using variadic parameters better performance
func IndexOf[T comparable](slice []T, element T, start *int) int {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return -1
	}

	startIndex := 0
	if start != nil {
		startIndex = convertIndexSimply(sliceLength, *start)
		if startIndex == -1 {
			return -1
		}
	}

	for i := startIndex; i < sliceLength; i++ {
		if slice[i] == element {
			return i // Short-circuit instantly
		}
	}

	return -1
}

// IndexOf method returns the first index at which a given element can be found in the array.
// Note: Prefer using IndexOf function for better performance
func (a Array[T]) IndexOf(element T, start *int) int {
	var (
		startIndex int = 0
		err        error
	)

	if start != nil {
		startIndex, err = ConvertIndex(a, *start, "start index")
		if err != nil {
			return -1
		}
	}

	sliceLength := len(a)

	if IsComparable(element) {
		for i := startIndex; i < sliceLength; i++ {
			if any(element) == any(a[i]) {
				return i
			}
		}
	} else {
		for i := startIndex; i < sliceLength; i++ {
			value := a[i]

			if reflect.DeepEqual(element, value) {
				return i
			}
		}
	}

	return -1
}
