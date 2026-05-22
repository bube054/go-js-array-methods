package array

// Map creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func Map[T any, V any](slice []T, fn MapFunc[T, V]) []V {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return make([]V, 0)
	}

	newList := make([]V, sliceLen)

	for i := 0; i < sliceLen; i++ {
		newList[i] = fn(slice[i], i, slice)
	}

	return newList
}

// Map creates a new slice from calling a function for every slice element. Map creates a new slice of the original slice type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func MapStrict[T any](slice []T, fn MapFuncStrict[T]) []T {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return slice
	}

	newList := make([]T, sliceLen)

	for i := 0; i < sliceLen; i++ {
		newList[i] = fn(slice[i], i, slice)
	}

	return newList
}

// Map method creates a new array with the results of calling a function for every array element.
func (a Array[T]) Map(fn MapFunc[T, any]) Array[any] {
	return Map(a, fn)
}

// MapStrict method creates a new array of the same type with the results of calling a function for every array element.
func (a Array[T]) MapStrict(fn MapFuncStrict[T]) Array[T] {
	return MapStrict(a, fn)
}
