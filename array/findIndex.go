package array

// FindIndex returns the index of the first element that satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - A result of -1 means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) search with no heap allocations.
//
// Example:
//
//	numbers := []int{1, 4, 7}
//	index := FindIndex(numbers, func(value int, _ int, _ []int) bool {
//		return value > 5
//	})
//	// index == 2
func FindIndex[T any](slice []T, fn Predicate[T]) int {
	for index, val := range slice {
		if fn(val, index, slice) {
			return index
		}
	}

	return -1
}

// FindIndex returns the index of the first element in the receiver that satisfies the predicate.
// It is the receiver-based form of FindIndex.
func (a Array[T]) FindIndex(fn Predicate[T]) int {
	return FindIndex(a, fn)
}
