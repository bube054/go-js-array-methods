package array

func FindIndex[T comparable] (slice []T, fn Predicate[T]) int {
	for index, val := range slice {
		if fn(val, index, slice) {
			return index
		}
	}

	return -1
}