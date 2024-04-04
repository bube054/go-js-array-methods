package array

// The Unshift() function adds new elements to the beginning of an array. The Unshift() function does not overwrite the original array.
func UnShift[T comparable](slice []T, element... T) []T {
	newList := append([]T{}, element...)
	return append(newList, slice...)
}