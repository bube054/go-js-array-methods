package array

// Push returns a new slice with additional elements appended.
//
// Behavior:
//   - The original slice is not modified.
//   - The function appends to the returned slice and may grow its backing capacity.
//
// Performance Note:
//   - This is an append-based path; allocation only occurs when the destination capacity is exhausted.
//
// Example:
//
//	numbers := []int{1, 2}
//	result := Push(numbers, 3, 4)
//	// result == []int{1, 2, 3, 4}
func Push[T any](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// Push returns a new array with additional elements appended.
// It is the receiver-based form of Push.
func (a Array[T]) Push(elements ...T) Array[T] {
	return Push(a, elements...)
}
