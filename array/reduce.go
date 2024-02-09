package array

func Reduce[T comparable](s []T, f ReduceFunc[T], defaultValue any) any {
	var accumulator = defaultValue

	for index, value := range s {
		accumulator = f(accumulator, value, index, s)
	}

	return accumulator
}

func ReduceStrict[T comparable, K comparable](s []T, f ReduceStrictFunc[T, K], defaultValue K) K {
	var accumulator = defaultValue

	for index, value := range s {
		accumulator = f(accumulator, value, index, s)
	}

	return accumulator
}

// d := []string{"ant", "bison", "camel", "duck", "elephant"}

	// r := array.Reduce[string](d, func(acc any, e string, i int, s []string) any {
	// 	accX := acc.(string)
	// 	return accX + " " + e
	// }, "")

// r := array.ReduceStrict[string, string](d, func(acc, e string, i int, s []string) string {
// 	return acc + " " + e
// }, "")

// fmt.Println(r)
