package array

import "reflect"

// IndexOf returns the first index of the target value, or -1 when not found.
//
// Behavior:
//   - The search is left-to-right.
//   - Negative start offsets are normalized from the end of the slice.
//   - Non-comparable values fall back to reflect.DeepEqual.
//
// Performance Note:
//   - The comparable fast path is allocation-free and performs a bare-metal loop.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	index := IndexOf(numbers, 2, nil)
//	// index == 1
func IndexOf[T comparable](slice []T, element T, start *int) int {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return -1
	}

	startIndex := 0
	if start != nil {
		startIndex = convertIndex(sliceLength, *start)
		if startIndex == -1 {
			return -1
		}
	}

	for i := startIndex; i < sliceLength; i++ {
		if slice[i] == element {
			return i
		}
	}

	return -1
}

// IndexOf returns the first index of the target value in the receiver, or -1 when not found.
// It is the receiver-based form of IndexOf.
func (a Array[T]) IndexOf(element T, start *int) int {
	startIndex := 0

	if start != nil {
		startIndex = convertIndex(len(a), *start)
		if startIndex == -1 {
			return -1
		}
	}

	sliceLength := len(a)

	if isComparable(element) {
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
