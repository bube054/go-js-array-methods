package array

// Every returns true when every element satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - The function short-circuits on the first false result.
//
// Performance Note:
//   - This is a bare-metal O(n) traversal with no heap allocations.
//
// Example:
//
//	numbers := []int{2, 4, 6}
//	allEven := Every(numbers, func(value int, _ int, _ []int) bool {
//		return value%2 == 0
//	})
//	// allEven == true
func Every[T any](slice []T, fn Predicate[T]) bool {
	for i := 0; i < len(slice); i++ {
		if !fn(slice[i], i, slice) {
			return false
		}
	}

	return true
}

// Every returns true when every element in the receiver satisfies the predicate.
// It is the receiver-based form of Every.
func (a Array[T]) Every(fn Predicate[T]) bool {
	return Every(a, fn)
}
