package array

// Shift function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element.
//
// Performance Note for Future Architect:
// Shift currently returns a pointer to the removed element (*T), which forces
// the Go compiler's escape analysis to push the element onto the heap (1 alloc/op).
// Additionally, because the pointer looks directly into the removed element, the
// underlying backing array slot cannot be zeroed out without corrupting caller data.
//
// In a future breaking release (v2.0), this signature should be changed to
// return by value (slice []T, element T, ok bool) to eliminate heap allocations
// completely and allow proper garbage collection cleanup of the shifted slot.
func Shift[T any](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	firstElement := slice[0]

	return slice[1:], &firstElement
}

// Shift method removes the first element from the array and returns the modified array and the removed element.
func (a Array[T]) Shift() (Array[T], *T) {

	return Shift(a)
}
