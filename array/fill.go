package array

import (
	"fmt"
)

// The Fill() function fills specified elements in an slice with a value. The fill() method does not overwrites the original slice.
func Fill[T comparable](slice []T, value T, start int, end int) ([]T, error) {
	sliceLength := len(slice)
	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newStartIndex, err := ConvertIndex(slice, start, "start index")

	if err != nil {
		return nil, err
	}

	newEndIndex, err := ConvertIndex(slice, end, "end index")

	if err != nil {
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