package array

import "fmt"

// At function returns an indexed element from a slice and returns a possible error relating to out of range indexes.
func At[T any](slice []T, index int) (element T, err error) {
	newIndex := convertIndex(len(slice), index)

	if newIndex == -1 {
		return element, fmt.Errorf("index: %d out of range", index)
	}

	return slice[newIndex], err
}

// At method returns an indexed element from the array and returns a possible error relating to out of range indexes.
func (a Array[T]) At(index int) (T, error) {
	return At(a, index)
}
