package array

func ForEach[S ~[]T, T any](slice S, fn ForEachFunc[S, T]) {
	for index, value := range slice {
		fn(value, index, slice)
	}
}

// The ForEach() method executes a provided function once for each array element.
func (a Array[T]) ForEach(fn ForEachFunc[Array[T], T]) {
	ForEach(a, fn)
}
