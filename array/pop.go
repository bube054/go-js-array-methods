package array

// The Pop() function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element.
func Pop[T comparable](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	lastElement := slice[sliceLength-1]

	return slice[:sliceLength-1], &lastElement
}