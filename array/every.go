package array

// The Every() function executes a function for each slice element. The Every() function returns true if the function returns true for all elements. The Every() function returns false if the function returns false for one element. The Every() function does not execute the function for empty elements. The Every() function does not change the original slice.
func Every[T comparable](slice []T, fn Predicate[T]) bool {
	for index, value := range slice {
		ok := fn(value, index, slice)
		if !ok {
			return false
		}
	}

	return true
}

// The Every() method tests whether all elements in the array pass the provided predicate function.
func (a Array[T]) Every(fn Predicate[T]) bool {
	return Every([]T(a), fn)
}
