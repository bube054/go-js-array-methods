package array

// ToString returns a comma-separated string representation of the slice.
//
// Behavior:
//   - The original slice is not modified.
//   - This is a thin wrapper around Join with the default separator.
//
// Performance Note:
//   - The cost is dominated by Join; for non-string slices, string conversion is performed during formatting.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := ToString(numbers)
//	// result == "1,2,3"
func ToString[T any](slice []T) string {
	return Join(slice, nil)
}

// ToString returns a comma-separated string representation of the receiver.
// It is the receiver-based form of ToString.
func (a Array[T]) ToString() string {
	return a.Join()
}
