package array

import (
	"errors"
	"fmt"
	"math"
)

func At[T comparable](s []T, i int) (*T, error) {
	if math.Abs(float64(i)) > float64(len(s)-1) {
		return nil, errors.New(fmt.Sprintf("%v is out of range", i))
	}

	index := i
	if index < 0 {
		index = len(s) + index
	}

	return &s[index], nil
}