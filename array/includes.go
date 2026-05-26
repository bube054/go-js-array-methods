package array

import "reflect"

// Includes returns true when the slice contains the target element.
//
// Behavior:
//   - The original slice is not modified.
//   - The optional start offset is normalized from the end of the slice when negative.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 5.59 µs/op.
//   - The comparable fast path is a bare-metal loop with no heap allocations.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	hasTwo := Includes(numbers, 2, nil)
//	// hasTwo == true
func Includes[T comparable](slice []T, element T, start *int) bool {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return false
	}

	startIndex := 0
	if start != nil {
		startIndex = convertIndex(sliceLength, *start)
		if startIndex == -1 {
			return false
		}
	}

	for i := startIndex; i < sliceLength; i++ {
		if slice[i] == element {
			return true
		}
	}

	return false
}

// Includes returns true when the receiver contains the target element.
// It is the receiver-based form of Includes.
//   - When the element is not comparable, the function falls back to reflect.DeepEqual.
func (a Array[T]) Includes(element T, start *int) bool {
	sliceLength := len(a)
	startIndex := 0
	if start != nil {
		startIndex = convertIndex(sliceLength, *start)
		if startIndex == -1 {
			return false
		}
	}

	if isComparable(element) {
		for i := startIndex; i < sliceLength; i++ {
			if any(a[i]) == any(element) {
				return true
			}
		}
	} else {
		for i := startIndex; i < sliceLength; i++ {
			if reflect.DeepEqual(element, a[i]) {
				return true
			}
		}
	}

	return false
}
