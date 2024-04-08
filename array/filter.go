package array

// The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements.
func Filter[T comparable](slice []T, fn Predicate[T]) []T {
	newList := []T{}

	for index, value := range slice {
		ok := fn(value, index, slice)
		if ok {
			newList = append(newList, value)
		}
	}

	return newList
}