package array

// Some returns true when at least one element satisfies the predicate.
//
// Behavior:
//   - The original slice is not modified.
//   - The predicate is not called for empty slices.
//   - The function short-circuits on the first true result.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 2.09 ns/op.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	hasEven := Some(numbers, func(value int, _ int, _ []int) bool {
//		return value%2 == 0
//	})
//	// hasEven == true
func Some[T any](slice []T, fn Predicate[T]) bool {
	sliceLen := len(slice)

	for i := 0; i < sliceLen; i++ {
		if fn(slice[i], i, slice) {
			return true
		}
	}

	return false
}

// Some returns true when at least one element in the receiver satisfies the predicate.
// It is the receiver-based form of Some.
func (a Array[T]) Some(fn Predicate[T]) bool {
	return Some(a, fn)
}
