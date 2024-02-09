package array

func Reverse[T comparable](s []T) []T {
	newList := []T{}

	for i := len(s) - 1; i >= 0; i-- {
		item := s[i]
		newList = append(newList, item)
	}

	return newList
}