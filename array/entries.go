package array

type Entry[T comparable] struct {
	index int
	element T
}

// The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice.
func Entries[T comparable](slice []T) []Entry[T] {
	newSlice := make([]Entry[T], len(slice))

	for i := range slice {
		element := slice[i]
		newSlice[i] = Entry[T]{index: i, element: element}
	}

	return newSlice
}