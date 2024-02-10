package array

import (
	"errors"
)

func Fill[T comparable](s []T, el T, startPos, endPos uint) ([]T, error) {
	length := len(s)

	if length == 0 {
		return nil, nil
		
	}

	if startPos < 0 || startPos >= uint(length) {
		return nil, errors.New("start position is out of range")
	}

	if endPos < startPos || endPos >= uint(length) {
		return nil, errors.New("end position is out of range")
	}

	newList := make([]T, length)
	copy(newList, s)

	for i := startPos; i <= endPos; i++ {
		newList[i] = el
	}

	return newList, nil
}
