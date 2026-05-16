package array

// Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func Map[S ~[]T, T any, R ~[]V, V any](slice S, fn MapFunc[S, T, V]) R {
	newList := make([]V, len(slice))

	for index, value := range slice {
		newItem := fn(value, index, slice)
		newList[index] = newItem
	}

	return newList
}

// Map() creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func MapStrict[S ~[]T, T any](slice S, fn MapFuncStrict[S, T]) S {
	newList := make([]T, len(slice))

	for index, value := range slice {
		newItem := fn(value, index, slice)
		newList[index] = newItem
	}

	return newList
}

// The Map() method creates a new array with the results of calling a function for every array element.
func (a Array[T]) Map(fn MapFunc[Array[T], T, any]) Array[any] {
	return Map[Array[T], T, Array[any]](a, fn)
}

// The MapStrict() method creates a new array of the same type with the results of calling a function for every array element.
func (a Array[T]) MapStrict(fn MapFuncStrict[Array[T], T]) Array[T] {
	return MapStrict(a, fn)
}
