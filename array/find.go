package array

func Find[T comparable](slice []T, fn Predicate[T]) (*T) {
	for index, val := range slice {
		if fn(val, index, slice) {
			return &val
		}
	}

	return nil
}
