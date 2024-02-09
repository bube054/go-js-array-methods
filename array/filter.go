package array

func Filter[T comparable](s []T, f Predicate[T]) []T {
	newList := []T{}

	for index, value := range s {
		ok := f(value, index, s)
		if ok {
			newList = append(newList, value)
		}
	}

	return newList
}

	// r := array.Filter[string](d, func(e string, i int, s []string) bool {
	// 	if len(e) > 4 {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// })

	// fmt.Println(r)