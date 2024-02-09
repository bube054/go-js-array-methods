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

// a := []int{1,2,3}
// r := array.Every[int](a, func(e int, i int, s []int) bool {
// 	return e < 2
// })
