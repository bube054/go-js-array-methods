package array

// ForEach invokes the callback once for each element in the slice.
//
// Behavior:
//   - The original slice is not modified.
//   - The callback runs in index order.
//
// Performance Note:
//   - This is a bare-metal O(n) traversal with no heap allocations.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	ForEach(numbers, func(value int, _ int, _ []int) {
//		_ = value
//	})
func ForEach[T any](slice []T, fn ForEachFunc[T]) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i], i, slice)
	}
}

// ForEach invokes the callback once for each element in the receiver.
// It is the receiver-based form of ForEach.
func (a Array[T]) ForEach(fn ForEachFunc[T]) {
	ForEach(a, fn)
}
