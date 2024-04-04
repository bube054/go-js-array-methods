package array

import (
	"fmt"
)

// The Slice() function returns selected elements in an slice, as a new slice and and error. The Slice() function selects from a given start, up to a (not inclusive) given end. The Slice() function does not change the original slice.
func Slice[T comparable](slice []T, start, end int) ([]T, error) {
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

	slicePartition := slice[newStartIndex:newEndIndex]

	return slicePartition, nil
}