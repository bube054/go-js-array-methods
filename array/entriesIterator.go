package array

import "iter"

// EntriesIterator returns a lazy, zero-allocation iterator over the slice elements.
//
// Usage:
//
//	for index, element := range arrays.EntriesIterator(mySlice) { ... }
func EntriesIterator[T any](slice []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := 0; i < len(slice); i++ {
			// If the consumer breaks out of the loop early, yield returns false
			if !yield(i, slice[i]) {
				return
			}
		}
	}
}

func (a Array[T]) EntriesIterator() iter.Seq2[int, T] {
	return EntriesIterator(a)
}
