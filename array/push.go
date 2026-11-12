package array

// Push returns a new slice with additional elements appended.
//
// Behavior:
//   - The original slice is not modified, even if it has spare capacity.
//   - The result always has its own backing array.
//
// Performance Note:
//   - Push copies the slice once (O(n)). A plain append would write into the
//     input's spare capacity, so two Pushes on the same slice, or a Push after
//     Pop, could overwrite each other's or the original's elements.
//
// Example:
//
//	numbers := []int{1, 2}
//	result := Push(numbers, 3, 4)
//	// result == []int{1, 2, 3, 4}
func Push[T any](slice []T, elements ...T) []T {
	result := make([]T, len(slice)+len(elements))
	copy(result, slice)
	copy(result[len(slice):], elements)
	return result
}

// Push returns a new array with additional elements appended.
// It is the receiver-based form of Push.
func (a Array[T]) Push(elements ...T) Array[T] {
	return Push(a, elements...)
}
