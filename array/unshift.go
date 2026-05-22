package array

// Unshift function adds new elements to the beginning of an slice. The Unshift() function does not overwrite the original slice.
func UnShift[T any](slice []T, element ...T) []T {
	elementLen := len(element)
	sliceLen := len(slice)
	finalLen := elementLen + sliceLen

	if finalLen == 0 {
		return slice
	}

	newList := make([]T, finalLen)

	copy(newList[:elementLen], element)
	copy(newList[elementLen:], slice)

	return newList
}

// UnShift method adds one or more elements to the beginning of the array and returns the new array.
func (a Array[T]) UnShift(elements ...T) Array[T] {
	return UnShift(a, elements...)
}
