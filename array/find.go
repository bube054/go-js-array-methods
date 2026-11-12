package array

// Find returns a pointer to the first element that satisfies the predicate.
//
// Behavior:
//   - The function itself never modifies the slice.
//   - The predicate is not called for empty slices.
//   - A nil result means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) search with no heap allocations.
//
// Pointer semantics:
//   - The result points at the element inside your slice, not at a copy.
//     Writing through it updates that element: use this to change a match in
//     place. For a copy, dereference it: value := *match.
//   - If the slice is later reallocated (for example by append), the pointer
//     keeps referring to the old backing array.
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
