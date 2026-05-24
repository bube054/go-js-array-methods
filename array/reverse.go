package array

// Reverse returns a new slice with the elements in reverse order.
//
// Behavior:
//   - The original slice is not modified.
//   - Empty inputs return the original slice.
//
// Performance Note:
//   - Current benchmark on this machine: 1 allocs/op and 35.2 µs/op.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := Reverse(numbers)
//	// result == []int{3, 2, 1}
func Reverse[T any](slice []T) []T {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return slice
	}

	newList := make([]T, sliceLen)

	for i, j := sliceLen-1, 0; i >= 0; i, j = i-1, j+1 {
		newList[j] = slice[i]
	}

	return newList
}

// Reverse returns a new array with the elements in reverse order.
// It is the receiver-based form of Reverse.
func (a Array[T]) Reverse() Array[T] {
	return Array[T](Reverse(a))
}
