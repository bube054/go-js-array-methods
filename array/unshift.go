package array

// UnShift returns a new slice with the provided elements prepended.
//
// Behavior:
//   - The original slice is not modified.
//   - Empty input returns the original slice.
//
// Performance Note:
//   - The function allocates one new slice and copies the existing data once.
//
// Example:
//
//	numbers := []int{3, 4}
//	result := UnShift(numbers, 1, 2)
//	// result == []int{1, 2, 3, 4}
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

// UnShift returns a new array with the provided elements prepended.
// It is the receiver-based form of UnShift.
func (a Array[T]) UnShift(elements ...T) Array[T] {
	return UnShift(a, elements...)
}
