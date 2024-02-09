package array

func UnShift[T comparable](s []T, e T) []T {
	newList := []T{e}
	return append(newList, s...)
}

func UnShiftMut[T comparable](s *[]T, e T) {
	newList := []T{e}
	*s = append(newList, (*s)...)
}
