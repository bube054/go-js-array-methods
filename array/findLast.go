package array

// The FindLast() function returns the value of the last element that passes a test. The FindLast() function executes a function for each slice element. The FindLast() function returns -1 if no elements are found. The FindLast() function does not execute the function for empty elements. The FindLast() function does not change the original slice.
func FindLast[T comparable](slice []T, fn Predicate[T]) *T {
	for i := len(slice) - 1; i >= 0; i-- {
		val := slice[i]
		if fn(val, i, slice) {
			return &val
		}
	}

	return nil
}
