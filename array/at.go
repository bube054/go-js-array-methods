package array

import "fmt"

// At returns the element at index or an error when the index is out of range.
//
// Behavior:
//   - Negative indexes are normalized from the end of the slice.
//   - Out-of-range indexes return an error.
//
// Performance Note:
//   - The common path is allocation-free and performs a constant-time lookup.
//
// Example:
//
//	numbers := []int{10, 20, 30}
//	value, err := At(numbers, 1)
//	// value == 20
//	// err == nil
func At[T any](slice []T, index int) (element T, err error) {
	newIndex := convertIndex(len(slice), index)

	if newIndex == -1 {
		return element, fmt.Errorf("index: %d out of range", index)
	}

	return slice[newIndex], err
}

// At returns the element at index or an error when the index is out of range.
// It is the receiver-based form of At.
func (a Array[T]) At(index int) (T, error) {
	return At(a, index)
}
