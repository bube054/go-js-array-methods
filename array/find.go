package array

// Find returns a pointer to the first element that satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - A nil result means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) search with no heap allocations.
//
// Warning:
//   - The returned pointer aliases the input slice. Do not mutate or re-slice the source while retaining the pointer.
//
// Example:
//
//	numbers := []int{1, 4, 7}
//	match := Find(numbers, func(value int, _ int, _ []int) bool {
//		return value > 5
//	})
//	// match != nil
//	// *match == 7
func Find[T any](slice []T, fn Predicate[T]) *T {
	for i := 0; i < len(slice); i++ {
		if fn(slice[i], i, slice) {
			return &slice[i]
		}
	}

	return nil
}

// Find returns a pointer to the first element in the receiver that satisfies the predicate.
// It is the receiver-based form of Find.
func (a Array[T]) Find(fn Predicate[T]) *T {
	return Find(a, fn)
}
