package array

func Entries[T comparable](s []T) []Entry[T] {
	result := []Entry[T]{}

	for index, value := range s{
		result = append(result, Entry[T]{index: index, value: value})
	}

	return result
}	