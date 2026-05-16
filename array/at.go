package array

import "fmt"

// The At() function returns an indexed element from a slice and returns a possible error relating to out of range indexes.
func At[S ~[]T, T any](slice S, index int) (element T, err error) {
	newIndex := convertIndexSimply(len(slice), index)

	if newIndex == -1 {
		return element, fmt.Errorf("index: %d out of range", index)
	}

	return slice[newIndex], err
}

// The At() method returns an indexed element from the array and returns a possible error relating to out of range indexes.
func (a Array[T]) At(index int) (T, error) {
	return At(a, index)
}
