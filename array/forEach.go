package array

func ForEach[T comparable](s []T, f ForEachFunc[T]) {
	for index, value := range s {
		f(value, index, s)
	}
}
