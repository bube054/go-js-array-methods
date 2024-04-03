package array

// Map() creates a new slice from calling a function for every slice element. Map creates a new slice of any type. Map() does not execute the function for empty elements. Map() does not change the original slice.
func Map[T comparable](slice []T, fn MapFunc[T]) []any {
	newList := make([]any, len(slice))

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
