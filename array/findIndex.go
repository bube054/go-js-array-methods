package array

func FindIndex[T comparable] (s []T, f Predicate[T]) int {
	for index, val := range s {
		if f(val, index, s) {
			return index
		}
	}

	return -1
}