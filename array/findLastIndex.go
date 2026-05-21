package array

func FindLastIndex[T any](slice []T, fn Predicate[T]) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if fn(slice[i], i, slice) {
			return i
		}
	}

	return -1
}

// The FindLastIndex() method returns the index of the last element that passes the provided predicate function.
func (a Array[T]) FindLastIndex(fn Predicate[T]) int {
	return FindLastIndex(a, fn)
}
