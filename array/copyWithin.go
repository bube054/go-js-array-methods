package array

import "fmt"

// CopyWithin returns a new slice with values copied into the target position.
//
// Behavior:
//   - The original slice is not modified.
//   - Negative indexes are normalized from the end of the slice.
//   - Invalid bounds return an error.
//
// Performance Note:
//   - The function allocates one output slice and performs a single copy of the affected region.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	result, err := CopyWithin(numbers, 1, 2, 4)
//	// result == []int{1, 3, 4, 4}
//	// err == nil
func CopyWithin[T any](slice []T, target int, start int, end int) ([]T, error) {
	sliceLength := len(slice)

	newTargetIndex := convertIndex(sliceLength, target)
	if newTargetIndex == -1 {
		return nil, fmt.Errorf("target index: %d out of range", target)
	}

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

	rangeDifference := newEndIndex - newStartIndex
	if newTargetIndex+rangeDifference > sliceLength {
		rangeDifference = sliceLength - newTargetIndex
	}

	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	if rangeDifference <= 0 {
		return newSlice, nil
	}

	srcSubSlice := slice[newStartIndex : newStartIndex+rangeDifference]
	dstSubSlice := newSlice[newTargetIndex : newTargetIndex+rangeDifference]
	copy(dstSubSlice, srcSubSlice)

	return newSlice, nil
}

// CopyWithin returns a new array with values copied into the target position.
// It is the receiver-based form of CopyWithin.
func (a Array[T]) CopyWithin(target, start, end int) (Array[T], error) {
	result, err := CopyWithin(a, target, start, end)
	return Array[T](result), err
}
