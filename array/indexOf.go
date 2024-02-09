package array

func IndexOf[T comparable](s []T, e T) int {
	for index, val := range s {
		if val == e {
			return index
		}
	}

	return -1
}