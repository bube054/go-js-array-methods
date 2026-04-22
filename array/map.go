package array

// Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func Map[T comparable, V any](slice []T, fn MapFunc[T, V]) []V {
	newList := make([]V, len(slice))

	for index, value := range slice {
		newItem := fn(value, index, slice)
		newList[index] = newItem
	}

	return newList
}

// Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func MapStrict[T comparable](slice []T, fn MapFuncStrict[T]) []T {
	newList := make([]T, len(slice))

	for index, value := range slice {
		newItem := fn(value, index, slice)
		newList[index] = newItem
	}

	return newList
}

// The Map() method creates a new array with the results of calling a function for every array element.
func (a Array[T]) Map(fn MapFunc[T, any]) Array[any] {
	return Array[any](Map([]T(a), fn))
}

// The MapStrict() method creates a new array of the same type with the results of calling a function for every array element.
func (a Array[T]) MapStrict(fn MapFuncStrict[T]) Array[T] {
	return Array[T](MapStrict([]T(a), fn))
}
