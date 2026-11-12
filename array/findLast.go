package array

// FindLast returns a pointer to the last element that satisfies the predicate.
//
// Behavior:
//   - The function itself never modifies the slice.
//   - The predicate is not called for empty slices.
//   - A nil result means no element matched.
//
// Performance Note:
//   - This is a bare-metal O(n) reverse search with no heap allocations.
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
