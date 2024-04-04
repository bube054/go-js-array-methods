package array

// The Some() function checks if any slice elements pass a test (provided as a callback function). The Some() function executes the callback function once for each slice element.The Some() function returns true (and stops) if the function returns true for one of the slice elements. The Some() function returns false if the function returns false for all of the slice elements. The Some() function does not execute the function for empty slice elements. The Some() function does not change the original slice. 
func Some[T comparable](slice []T, fn Predicate[T]) bool {
	for index, value := range slice {
		ok := fn(value, index, slice)
		if ok {
			return true
		}
	}

	return false
}