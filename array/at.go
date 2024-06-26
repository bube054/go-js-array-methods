// Package array provides functions for manipulating golang slices.
package array

// The At() function returns an indexed element from a slice and returns a possible error relating to out of range indexes.
func At[T comparable](slice []T, index int) (T, error) {
	var elementAtPosition T

	newIndex, err := ConvertIndex(slice, index, "index")

	if err != nil {
		return elementAtPosition, err
	}

	elementAtPosition = slice[newIndex]

	return elementAtPosition, nil
}
