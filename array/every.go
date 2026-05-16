package array

// The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. The Every() function does not change the original slice.
func Every[S ~[]T, T any](slice S, fn Predicate[S, T]) bool {
	for i := 0; i < len(slice); i++ {
		if !fn(slice[i], i, slice) {
			return false
		}
	}

	return true
}

// The Every() method tests whether all elements in the array pass the provided predicate function.
func (a Array[T]) Every(fn Predicate[Array[T], T]) bool {
	return Every(a, fn)
}
