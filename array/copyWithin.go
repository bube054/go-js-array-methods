package array

import (
	"fmt"
	// "math"
)

// The CopyWithin() function returns a copy of slice elements to another position in an slice and a possible error relating to out of range indexes.. The CopyWithin() function overwrites the existing values in the new slice.
func CopyWithin[T comparable](slice []T, target int, start int, end int) ([]T, error) {
	sliceLength := len(slice)
	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newTargetIndex, err := ConvertIndex(slice, target, "target index")

	if err != nil {
		return nil, err
	}

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

	rangeDifference := newEndIndex - newStartIndex

	pointer := newStartIndex

	for i := newTargetIndex; i < newTargetIndex+rangeDifference; i++ {
		newSlice[i] = slice[pointer]
		pointer++
	}

	return newSlice, nil
}
