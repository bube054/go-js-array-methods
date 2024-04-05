package array

// The FindIndex() function executes a function for each slice element. The FindIndex() function returns the index (position) of the first element that passes a test. The FindIndex() function returns -1 if no match is found. The FindIndex() function does not execute the function for empty slice elements. The FindIndex() function does not change the original slice.
func FindIndex[T comparable] (slice []T, fn Predicate[T]) int {
	for index, val := range slice {
		if fn(val, index, slice) {
			return index
		}
	}

	return -1
}