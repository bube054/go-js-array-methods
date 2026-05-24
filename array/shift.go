package array

// Shift returns a new slice without the first element and the removed value.
//
// Behavior:
//   - The original slice is not modified.
//   - Empty slices return the original slice and nil.
//
// Performance Note:
//   - The current signature returns a pointer to the removed element, which can trigger escape analysis and heap allocation.
//   - Roadmap Notice: a future major release should return the value directly to remove this escape path.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	trimmed, removed := Shift(numbers)
//	// trimmed == []int{2, 3}
//	// removed != nil
//	// *removed == 1
func Shift[T any](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	firstElement := slice[0]

	return slice[1:], &firstElement
}

// Shift returns a new array without the first element and the removed value.
// It is the receiver-based form of Shift.
func (a Array[T]) Shift() (Array[T], *T) {

	return Shift(a)
}
