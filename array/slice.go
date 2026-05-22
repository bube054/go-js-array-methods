package array

import (
	"fmt"
)

// Slice function returns selected elements in an slice, as a new slice and and error. The Slice() function selects from a given start, up to a (not inclusive) given end. The Slice() function does not change the original slice.
//
// Performance Note for Future Architect:
// This function maintains its signature for API stability. In a future major
// version update, consider making a major shifts:
//
// Drop the 'error' return type path completely in favor of safe bounding
// clamping (e.g., clamping out-of-bound indexes to 0 or len), reducing
// unnecessary call-site error propagation and string allocations.
func Slice[T any](slice []T, start, end int) ([]T, error) {
	sliceLength := len(slice)
	newStartIndex := convertIndexSimply(sliceLength, start)

	if newStartIndex == -1 {
		return nil, fmt.Errorf("start index: %d is out of bounds for slice of length: %d", start, sliceLength)
	}

	newEndIndex := convertEndIndex(sliceLength, end)

	if newEndIndex == -1 {
		return nil, fmt.Errorf("end index: %d is out of bounds for slice of length: %d", end, sliceLength)
	}

	if newStartIndex > newEndIndex {
		return nil, fmt.Errorf("start index: %d cannot be greater than end index: %d", start, end)
	}

	return slice[newStartIndex:newEndIndex], nil
}

// Slice method returns a shallow copy of a portion of the array.
func (a Array[T]) Slice(start, end int) (Array[T], error) {
	sliced, err := Slice(a, start, end)
	return Array[T](sliced), err
}
