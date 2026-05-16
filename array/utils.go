package array

import (
	"fmt"
)

// ConvertIndex converts negative JS-style indices to absolute Go slice indices.
// It achieves zero-allocation on the happy path by deferring error construction.
func ConvertIndex[T any](slice []T, index int, nameOfIndex string) (int, error) {
	newIndex := convertIndexSimply(len(slice), index)

	if newIndex == -1 {
		return -1, fmt.Errorf("%s: %d out of range", nameOfIndex, index)
	}

	return newIndex, nil
}

func convertIndexSimply(sliceLen, index int) int {
	if index < 0 {
		index += sliceLen
	}
	if index < 0 || index >= sliceLen {
		return -1
	}
	return index
}

func OptionalParam[T any](params []T, defaultValue T) T {
	if len(params) > 0 {
		return params[0]
	}
	return defaultValue
}
