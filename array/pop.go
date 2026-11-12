package array

// Pop returns a new slice without the last element and the removed value.
//
// Behavior:
//   - The original slice is not modified.
//   - The result shares memory with the original but has no spare capacity, so
//     appending to it (including with Push) never overwrites the original. Use
//     slices.Clone if you will assign to its elements.
//   - Empty slices return the original slice and nil.
//
// Performance Note:
//   - The current signature returns a pointer to the removed element, which can trigger escape analysis and heap allocation.
//   - Roadmap Notice: a future major release should return the value directly to remove this escape path.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	trimmed, removed := Pop(numbers)
//	// trimmed == []int{1, 2}
//	// removed != nil
//	// *removed == 3
func Pop[T any](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	lastElement := slice[sliceLength-1]

	// The third index caps the capacity, so an append must allocate new memory
	// instead of writing over the removed last element of the original.
	return slice[: sliceLength-1 : sliceLength-1], &lastElement
}

// Pop returns a new array without the last element and the removed value.
// It is the receiver-based form of Pop.
func (a Array[T]) Pop() (Array[T], *T) {

	return Pop(a)
}
