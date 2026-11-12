package array

import "fmt"

// Slice returns a new slice containing the selected range.
//
// Behavior:
//   - The original slice is not modified.
//   - Negative indexes are normalized from the end of the slice.
//   - Invalid bounds return an error.
//
// Performance Note:
//   - 0 allocs/op: the function returns a subslice backed by the original array
//     without copying.
//   - The result's capacity ends at its last element, so appending to it
//     (including with Push) never overwrites the rest of the original. Use
//     slices.Clone if you will assign to its elements.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	result, err := Slice(numbers, 1, 3)
//	// result == []int{2, 3}
//	// err == nil
func Slice[T any](slice []T, start, end int) ([]T, error) {
	sliceLength := len(slice)
	newStartIndex := convertIndex(sliceLength, start)

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

	// The third index caps the capacity, so an append must allocate new memory
	// instead of writing over the elements after the range.
	return slice[newStartIndex:newEndIndex:newEndIndex], nil
}

// Slice returns a new array containing the selected range.
// It is the receiver-based form of Slice.
func (a Array[T]) Slice(start, end int) (Array[T], error) {
	sliced, err := Slice(a, start, end)
	return Array[T](sliced), err
}
