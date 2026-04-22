package array

// The Unshift() function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice.
func UnShift[T comparable](slice []T, element ...T) []T {
	newList := append([]T{}, element...)
	return append(newList, slice...)
}

// The UnShift() method adds one or more elements to the beginning of the array and returns the new array.
func (a Array[T]) UnShift(elements ...T) Array[T] {
	return Array[T](UnShift([]T(a), elements...))
}
