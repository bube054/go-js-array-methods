package array

import (
	"errors"
)

func FindLast[T comparable](s []T, f Predicate[T]) (*T, error) {
	length := len(s)

	if length == 0 {
		return nil, errors.New("slice is empty")
	}

	for i := length - 1; i >= 0; i-- {
		val := s[i]
		if f(val, i, s) {
			return &val, nil
		}
	}

	return nil, errors.New("could not find element in slice")
}
