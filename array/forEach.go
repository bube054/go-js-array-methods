package array

func ForEach[T comparable](slice []T, fn ForEachFunc[T]) {
	for index, value := range slice {
		fn(value, index, slice)
	}
}

// The ForEach() method executes a provided function once for each array element.
func (a Array[T]) ForEach(fn ForEachFunc[T]) {
	ForEach([]T(a), fn)
}
