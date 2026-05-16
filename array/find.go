package array

// The Find() function returns a pointer value of the first element that passes a test. The Find() function executes a function for each slice element. The Find() function returns undefined if no elements are found. The Find() function does not execute the function for empty elements. The Find() function does not change the original slice
func Find[S ~[]T, T any](slice S, fn Predicate[S, T]) *T {
	for index, val := range slice {
		if fn(val, index, slice) {
			return &slice[index]
		}
	}

	return nil
}

// The Find() method returns the first element that passes the provided predicate function.
func (a Array[T]) Find(fn Predicate[Array[T], T]) *T {
	return Find(a, fn)
}
