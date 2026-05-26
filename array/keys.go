package array

import "iter"

// Keys returns an iterator over the indexes of the slice.
// It yields only the integer index, matching JavaScript's Array.prototype.keys().
//
// Behavior:
//   - The original slice is not modified.
//   - The iterator yields indexes from 0 to len(slice)-1.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 0.01 µs/op.
//   - The comparable fast path is a bare-metal loop with no heap allocations.
//
// Example usage:
//
//	for index := range Keys([]string{"a", "b", "c"}) {
//	    fmt.Println(index)
//	}
func Keys[T any](slice []T) iter.Seq[int] {
	return func(yield func(int) bool) {
		sliceLen := len(slice)
		for i := 0; i < sliceLen; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Keys returns an iterator over the array's indexes.
func (a Array[T]) Keys() iter.Seq[int] {
	return Keys(a)
}
