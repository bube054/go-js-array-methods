package array

// The filter() function creates a new slice filled with elements that pass a test provided by a function. The filter() function does not execute the function for empty elements.
func Filter[S ~[]T, T any](slice S, fn Predicate[S, T]) S {
	newList := make(S, 0, len(slice))

	for i := 0; i < len(slice); i++ {
		if fn(slice[i], i, slice) {
			newList = append(newList, slice[i])
		}
	}

	return newList
}

// The Filter() method creates a new array with elements that pass the provided predicate function.
func (a Array[T]) Filter(fn Predicate[Array[T], T]) Array[T] {
	return Array[T](Filter([]T(a), fn))
}
