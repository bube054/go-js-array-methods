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

// The At() method returns an indexed element from the array and returns a possible error relating to out of range indexes.
func (a Array[T]) At(index int) (T, error) {
	return At([]T(a), index)
}
