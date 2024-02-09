package array

func FindIndex[T comparable] (s []T, f Predicate[T]) int {
	for index, val := range s {
		if f(val, index, s) {
			return index
		}
	}

	return -1
}

	// d := []string{"ant", "bison", "camel", "duck", "elephant"}

	// r := array.FindIndex[string](d, func(e string, i int, s []string) bool {
	// 	if e == "bisone" {
	// 		return true
	// 	} else {
	// 		return false
	// 	}
	// })

	// fmt.Println(r)