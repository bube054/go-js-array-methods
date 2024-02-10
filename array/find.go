package array

import (
	"errors"
)

func Find[T comparable](s []T, f Predicate[T]) (*T, error) {
	length := len(s)

	if length == 0 {
		return nil, errors.New("slice is empty")
	}

	for index, val := range s {
		if f(val, index, s) {
			return &val, nil
		}
	}

	return nil, errors.New("could not find element in slice")
}