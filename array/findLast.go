package array

// FindLast returns a pointer to the last element that satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - A nil result means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) reverse search with no heap allocations.
//
// Warning:
//   - The returned pointer aliases the input slice. Do not mutate or re-slice the source while retaining the pointer.
//
// Example:
//
//	numbers := []int{1, 4, 7}
//	match := FindLast(numbers, func(value int, _ int, _ []int) bool {
//		return value%2 == 0
//	})
//	// match != nil
//	// *match == 4
func FindLast[T any](slice []T, fn Predicate[T]) *T {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i], i, slice) {
			return &slice[i]
		}
	}

	return nil
}

// FindLast returns a pointer to the last element in the receiver that satisfies the predicate.
// It is the receiver-based form of FindLast.
func (a Array[T]) FindLast(fn Predicate[T]) *T {
	return FindLast(a, fn)
}
