package array

import (
	"errors"
	"math"
)

// 1,2,3,4
func Fill[T comparable](s []T, el T, startPos, endPos uint) ([]T, error) {
	length := len(s)

	newList := []T{}

	if length == 0 {
		return newList, nil
	}

	if math.Abs(float64(startPos)) > float64(length-1) {
		return []T{}, errors.New("start position is out of range")
	}

	if math.Abs(float64(endPos)) > float64(length-1) {
		endPos = uint(length)
	}

	ForEach[T](s, func(e T, i int, s []T) {
		if uint(i) >= startPos && uint(i) < endPos {
			newList = append(newList, el)
		} else {
			newList = append(newList, e)
		}
	})

	return newList, nil
}
