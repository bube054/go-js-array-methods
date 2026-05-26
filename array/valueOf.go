package array

// ValueOf returns the input slice unchanged.
//
// Behavior:
//   - The original slice is not modified.
//   - This is an identity function.
//
// Performance Note:
//   - This is allocation-free and returns the same backing slice.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := ValueOf(numbers)
//	// result == numbers
func ValueOf[T any](slice []T) []T {
	return slice
}

// ValueOf returns the receiver unchanged.
// It is the receiver-based form of ValueOf.
func (a Array[T]) ValueOf() Array[T] {
	return ValueOf(a)
}
