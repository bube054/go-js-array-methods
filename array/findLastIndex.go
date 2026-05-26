package array

// FindLastIndex returns the index of the last element that satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - A result of -1 means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) reverse search with no heap allocations.
//
// Example:
//
//	numbers := []int{1, 4, 7}
//	index := FindLastIndex(numbers, func(value int, _ int, _ []int) bool {
//		return value%2 == 0
//	})
//	// index == 1
func FindLastIndex[T any](slice []T, fn Predicate[T]) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i], i, slice) {
			return i
		}
	}

	return -1
}

// FindLastIndex returns the index of the last element in the receiver that satisfies the predicate.
// It is the receiver-based form of FindLastIndex.
func (a Array[T]) FindLastIndex(fn Predicate[T]) int {
	return FindLastIndex(a, fn)
}
