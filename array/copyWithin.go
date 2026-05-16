package array

import (
	"fmt"
	// "math"
)

// The CopyWithin() function returns a copy of slice elements to another position in an slice and a possible error relating to out of range indexes.. The CopyWithin() function overwrites the existing values in the new slice.
func CopyWithin[S ~[]T, T any](slice S, target int, start int, end int) (S, error) {
	sliceLength := len(slice)

	// 1. Validate and Normalize indices BEFORE any heap allocations
	newTargetIndex := convertIndexSimply(sliceLength, target)
	if newTargetIndex == -1 {
		return nil, fmt.Errorf("target index: %d out of range", target)
	}

	newStartIndex := convertIndexSimply(sliceLength, start)
	if newStartIndex == -1 {
		return nil, fmt.Errorf("start index: %d out of range", start)
	}

	// Normalize end index (end is exclusive, so it can legally equal sliceLength)
	newEndIndex := convertEndIndex(sliceLength, end)

	if newEndIndex == -1 {
		return nil, fmt.Errorf("end index: %d out of range", end)
	}

	if newStartIndex > newEndIndex {
		return nil, fmt.Errorf("start index: %d cannot be greater than end index: %d", start, end)
	}

	// 2. Calculate copy range and apply JS-style clipping to prevent panics
	rangeDifference := newEndIndex - newStartIndex
	if newTargetIndex+rangeDifference > sliceLength {
		rangeDifference = sliceLength - newTargetIndex
	}

	// 3. Allocate and duplicate the base array ONLY on verified happy path
	newSlice := make(S, sliceLength)
	copy(newSlice, slice)

	// Short-circuit if there's nothing to copy
	if rangeDifference <= 0 {
		return newSlice, nil
	}

	// 4. Blazing fast block-memory move using Go's native copy over sub-slices
	srcSubSlice := slice[newStartIndex : newStartIndex+rangeDifference]
	dstSubSlice := newSlice[newTargetIndex : newTargetIndex+rangeDifference]
	copy(dstSubSlice, srcSubSlice)

	return newSlice, nil
}

// The CopyWithin() method copies array elements to another position in the array and returns the modified array.
func (a Array[T]) CopyWithin(target, start, end int) (Array[T], error) {
	return CopyWithin(a, target, start, end)
}
