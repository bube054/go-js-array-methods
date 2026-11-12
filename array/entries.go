package array

// Entries returns an eager slice of Entry values containing the source indexes and elements.
//
// Behavior:
//   - The original slice is not modified.
//   - The returned slice has one Entry per input element.
//
// Performance Note:
//   - This function allocates an O(N) result slice up front.
//   - If you need a streaming path for very large datasets, use EntriesIterator, which evaluates lazily with 0 allocs/op.
//
// Example:
//
//	numbers := []int{10, 20}
//	entries := Entries(numbers)
//	// entries[0].Index == 0
//	// entries[0].Value == 10
func Entries[T any](slice []T) []Entry[T] {
	newSlice := make([]Entry[T], len(slice))
	sliceLen := len(slice)

	for i := 0; i < sliceLen; i++ {
		newSlice[i] = Entry[T]{
			Index: i,
			Value: slice[i],
		}
	}

	return newSlice
}

// Entries returns an eager slice of Entry values for the receiver.
// It is the receiver-based form of Entries.
func (a Array[T]) Entries() []Entry[T] {
	return Entries(a)
}
