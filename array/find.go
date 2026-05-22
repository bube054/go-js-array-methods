package array

// Find function returns a pointer value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice.
// Warning: returned *T points into the backing array of the provided slice. Do not mutate or re-slice/re-append the original slice while retaining the returned pointer.
func Find[T any](slice []T, fn Predicate[T]) *T {
	for i := 0; i < len(slice); i++ {
		if fn(slice[i], i, slice) {
			return &slice[i]
		}
	}

	return nil
}

// Find method returns the first element that passes the provided predicate function.
func (a Array[T]) Find(fn Predicate[T]) *T {
	return Find(a, fn)
}
