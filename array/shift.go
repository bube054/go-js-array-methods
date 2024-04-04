package array

// The Shift() function removes the first item of an slice. The Shift() function changes the original slice. The Shift() function returns the new slice and a pointer to the  shifted element.
func Shift[T comparable](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	firstElement := slice[0]

	return slice[1:], &firstElement
}