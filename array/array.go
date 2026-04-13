package array

type Array[T comparable] []T

// The At() function returns an indexed element from a slice and returns a possible error relating to out of range indexes.
func (a *Array[T]) At(index int) (T, error) {
	return At([]T(*a), index)
}

