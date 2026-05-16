package array

// The Shift() function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element.
func Shift[S ~[]T, T any](slice S) (S, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	firstElement := slice[0]

	return slice[1:], &firstElement
}

// The Shift() method removes the first element from the array and returns the modified array and the removed element.
func (a Array[T]) Shift() (Array[T], *T) {

	return Shift(a)
}
