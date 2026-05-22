package array

import (
	"fmt"
	// "math"
)

// CopyWithin function returns a copy of slice elements to another position in an slice and a possible error relating to out of range indexes.. The CopyWithin() function overwrites the existing values in the new slice.
func CopyWithin[T any](slice []T, target int, start int, end int) ([]T, error) {
	sliceLength := len(slice)

	// 1. Validate and Normalize indices BEFORE any heap allocations
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

// CopyWithin method copies array elements to another position in the array and returns the modified array.
func (a Array[T]) CopyWithin(target, start, end int) (Array[T], error) {
	result, err := CopyWithin(a, target, start, end)
	return Array[T](result), err
}
