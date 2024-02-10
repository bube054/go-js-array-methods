package array

func FindLastIndex[T comparable](s []T, f Predicate[T]) int {
	length := len(s)

	for i := length - 1; i >= 0; i-- {
		val := s[i]
		if f(val, i, s) {
			return i
		}
	}

	return -1
}
