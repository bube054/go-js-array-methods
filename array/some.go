package array

func Some[T comparable](s []T, f Predicate[T]) bool {
	for index, value := range s {
		ok := f(value, index, s)
		if ok {
			return true
		}
	}

	return false
}

// a := []int{1,2,3}
// r := array.Some[int](a, func(e int, i int, s []int) bool {
// return e == 4
// })