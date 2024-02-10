package array

func Flat[T comparable](s [][]T) []T  {
	newList := []T{}

	for _, value := range s {
		for _, value2 := range value {
			newList = append(newList, value2)
		}
	}

	return newList
}