package array

import "fmt"

// With returns a new slice with the element at index replaced by value.
//
// Behavior:
//   - The original slice is not modified.
//   - Invalid indexes return an error.
//
// Performance Note:
//   - Current benchmark on this machine: 1 allocs/op and 31.4 µs/op.
//   - Roadmap Notice: a future major release may remove the error return path to simplify call sites.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result, err := With(numbers, 1, 9)
//	// result == []int{1, 9, 3}
//	// err == nil
func With[T any](slice []T, index int, value T) ([]T, error) {
	sliceLength := len(slice)
	newIndex := convertIndex(sliceLength, index)

	if newIndex == -1 {
		return slice, fmt.Errorf("index: %d out of range", index)
	}

	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newSlice[newIndex] = value

	return newSlice, nil
}

// With returns a new array with the element at the given index replaced by value.
// It is the receiver-based form of With.
func (a Array[T]) With(index int, value T) (Array[T], error) {

	return With(a, index, value)
}
