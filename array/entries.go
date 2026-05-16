package array

// The entries() function returns an Slice with structs with field/value: {index 0, element "Banana"}. The entries() function does not change the original slice.
//
// Performance Note for future scaling:
// This function evaluates eagerly, allocating an O(N) memory block upfront.
// If processing ultra-large datasets where garbage collection pressure becomes
// a bottleneck, consider using entriesIterator which is using Go 1.23+ range-over-func iterators
// (iter.Seq2[int, T]) to stream entries lazily with zero heap allocations.
func Entries[S ~[]T, T any](slice S) []Entry[T] {
	newSlice := make([]Entry[T], len(slice))
	sliceLen := len(slice)

	for i := 0; i < sliceLen; i++ {
		newSlice[i] = Entry[T]{
			index:   i,
			element: slice[i],
		}
	}

	return newSlice
}

// The Entries() method returns an array of [index, element] pairs for each element in the array.
func (a Array[T]) Entries() []Entry[T] {
	return Entries(a)
}
