package array

import (
	"fmt"
)

// The Fill() function returns fills of specified elements in an slice with a value and a possible error due to indexes out of range.
func Fill[S ~[]T, T any](slice S, value T, start int, end int) (S, error) {
	sliceLength := len(slice)
	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newStartIndex, err := ConvertIndex(slice, start, "start index")

	if err != nil {
		return nil, err
	}

	newEndIndex, err := ConvertIndex(slice, end, "end index")

	if err != nil {
		// newEndIndex = sliceLength
		return nil, err
	}

	if newStartIndex > newEndIndex {
		return nil, fmt.Errorf("start index: %d cannot be greater than end index: %d", start, end)
	}

	for i := newStartIndex; i < newEndIndex; i++ {
		newSlice[i] = value
	}

	return newSlice, nil
}

// The Fill() method fills array elements from a start index to an end index with a static value.
func (a Array[T]) Fill(value T, start, end int) (Array[T], error) {
	return Fill(a, value, start, end)
}
