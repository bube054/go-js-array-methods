package array

// The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements.
func Filter[S ~[]T, T any](slice S, fn Predicate[S, T]) S {
	newList := make([]T, 0)

	for index, value := range slice {
		ok := fn(value, index, slice)
		if ok {
			newList = append(newList, value)
		}
	}

	return newList
}

// The Filter() method creates a new array with elements that pass the provided predicate function.
func (a Array[T]) Filter(fn Predicate[Array[T], T]) Array[T] {
	return Array[T](Filter([]T(a), fn))
}
