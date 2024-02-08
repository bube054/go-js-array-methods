package array

func Push[T comparable](s []T, e T) []T {
	return append(s, e)
}

func PushMut[T comparable](s *[]T, e T) {
	*s = append(*s, e)
}
