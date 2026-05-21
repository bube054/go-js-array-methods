package array

// Unshift function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice.
func UnShift[T any](slice []T, element ...T) []T {
	newList := make([]T, 0, len(slice)+len(element))
	newList = append(newList, element...)
	return append(newList, slice...)
}

// UnShift method adds one or more elements to the beginning of the array and returns the new array.
func (a Array[T]) UnShift(elements ...T) Array[T] {
	return UnShift(a, elements...)
}
