package array

func Every[T comparable](s []T, f Predicate[T]) bool {
	for index, value := range s {
		ok := f(value, index, s)
		if !ok {
			return false
		}
	}

	return true
}