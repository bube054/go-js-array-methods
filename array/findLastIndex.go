package array

func FindLastIndex[T comparable](slice []T, fn Predicate[T]) int {
	for i := len(slice) - 1; i >= 0; i-- {
		val := slice[i]
		if fn(val, i, slice) {
			return i
		}
	}

	return -1
}
