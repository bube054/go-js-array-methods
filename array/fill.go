package array

import "fmt"

// Fill returns a new slice with the selected range replaced by value.
//
// Behavior:
//   - The original slice is not modified.
//   - Negative indexes are normalized from the end of the slice.
//   - Invalid bounds return an error.
//
// Performance Note:
//   - The function copies the slice once and mutates the output buffer in place.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	result, err := Fill(numbers, 9, 1, 3)
//	// result == []int{1, 9, 9, 4}
//	// err == nil
func Fill[T any](slice []T, value T, start int, end int) ([]T, error) {
	sliceLength := len(slice)

	newStartIndex := convertIndex(sliceLength, start)
	if newStartIndex == -1 {
		return nil, fmt.Errorf("start index: %d out of range", start)
	}

	newEndIndex := convertEndIndex(sliceLength, end)

	if newEndIndex == -1 {
		return nil, fmt.Errorf("end index: %d out of range", end)
	}

	if newStartIndex > newEndIndex {
		return nil, fmt.Errorf("start index: %d cannot be greater than end index: %d", start, end)
	}

	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	targetZone := newSlice[newStartIndex:newEndIndex]
	for i := range targetZone {
		targetZone[i] = value
	}

	return newSlice, nil
}

// Fill returns a new array with the selected range replaced by value.
// It is the receiver-based form of Fill.
func (a Array[T]) Fill(value T, start, end int) (Array[T], error) {
	result, err := Fill(a, value, start, end)

	return Array[T](result), err
}
