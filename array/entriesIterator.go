package array

import "iter"

// EntriesIterator returns a lazy, zero-allocation iterator over the slice elements.
//
// Usage:
//
//	for index, element := range arrays.EntriesIterator(mySlice) { ... }
func EntriesIterator[S ~[]T, T any](slice S) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i, v := range slice {
			// If the consumer breaks out of the loop early, yield returns false
			if !yield(i, v) {
				return
			}
		}
	}
}

func (a Array[T]) EntriesIterator() iter.Seq2[int, T] {
	return EntriesIterator(a)
}
