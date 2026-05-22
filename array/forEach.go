package array

func ForEach[T any](slice []T, fn ForEachFunc[T]) {
	for i := 0; i < len(slice); i++ {
		fn(slice[i], i, slice)
	}
}

// The ForEach() method executes a provided function once for each array element.
func (a Array[T]) ForEach(fn ForEachFunc[T]) {
	ForEach(a, fn)
}
