package array

import (
	"fmt"
)

func ConvertIndex[T comparable](slice []T, index int, nameOfIndex string) (int, error) {
	sliceLength := len(slice)
	err := fmt.Errorf("%v: %d out of range", nameOfIndex, index)

	if index > sliceLength-1 {
		return -1, err
	} else if index < 0 && index < -(sliceLength) {
		return -1, err
	} else if index >= 0 {
		return index, nil
	} else if index < 0 {
		return sliceLength + index, nil
	} else {
		return -1, err
	}
}

func OptionalParam[T any](params []T, defaultValue T) T {
	if len(params) > 0 {
		return params[0]
	}
	return defaultValue
}
