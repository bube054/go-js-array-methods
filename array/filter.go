package array

// filter function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements.
func Filter[T any](slice []T, fn Predicate[T]) []T {
	newList := make([]T, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if fn(slice[i], i, slice) {
			newList = append(newList, slice[i])
		}
	}

	return newList
}

// Filter method creates a new array with elements that pass the provided predicate function.
func (a Array[T]) Filter(fn Predicate[T]) Array[T] {
	return Array[T](Filter(a, fn))
}
