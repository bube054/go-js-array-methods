package array

import (
	"fmt"
)

// The Fill() function returns fills of specified elements in an slice with a value and a possible error due to indexes out of range.
func Fill[S ~[]T, T any](slice S, value T, start int, end int) (S, error) {
	sliceLength := len(slice)

	newStartIndex := convertIndexSimply(sliceLength, start)
	if newStartIndex == -1 {
		return nil, fmt.Errorf("start index: %d out of range", start)
	}

	// Handle exclusive end index logic manually to safely allow end == sliceLength
	newEndIndex := convertEndIndex(sliceLength, end)

	if newEndIndex == -1 {
		return nil, fmt.Errorf("end index: %d out of range", end)
	}

	if newStartIndex > newEndIndex {
		return nil, fmt.Errorf("start index: %d cannot be greater than end index: %d", start, end)
	}

	newSlice := make(S, sliceLength)
	copy(newSlice, slice)

	targetZone := newSlice[newStartIndex:newEndIndex]
	for i := range targetZone {
		targetZone[i] = value
	}

	return newSlice, nil
}

// The Fill() method fills array elements from a start index to an end index with a static value.
func (a Array[T]) Fill(value T, start, end int) (Array[T], error) {
	return Fill(a, value, start, end)
}
