package array

import "reflect"

// LastIndexOf returns the last index of the target value, or -1 when not found.
//
// Behavior:
//   - The search is right-to-left.
//   - Negative start offsets are normalized from the end of the slice.
//   - Non-comparable values fall back to reflect.DeepEqual.
//
// Performance Note:
//   - The comparable fast path is allocation-free and performs a bare-metal loop.
//
// Example:
//
//	numbers := []int{1, 2, 3, 2}
//	index := LastIndexOf(numbers, 2, nil)
//	// index == 3
func LastIndexOf[T comparable](slice []T, element T, start *int) int {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return -1
	}

	startIndex := sliceLength - 1
	if start != nil {
		startIndex = convertIndex(sliceLength, *start)

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

// LastIndexOf returns the last index of the target value in the receiver, or -1 when not found.
// It is the receiver-based form of LastIndexOf.
func (a Array[T]) LastIndexOf(element T, start *int) int {
	sliceLength := len(a)

	startIndex := sliceLength - 1

	if start != nil {
		startIndex = convertIndex(sliceLength, *start)
		if startIndex == -1 {
			return -1
		}
	}

	if isComparable(element) {
		for i := startIndex; i >= 0; i-- {
			if any(element) == any(a[i]) {
				return i
			}
		}
	} else {
		for i := startIndex; i >= 0; i-- {
			value := a[i]

			if reflect.DeepEqual(element, value) {
				return i
			}
		}
	}
	return -1
}
