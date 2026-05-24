package array

// Map returns a new slice containing the transformed values of the input.
//
// Behavior:
//   - The original slice is not modified.
//   - Empty inputs return an empty slice.
//
// Performance Note:
//   - Current benchmark on this machine: 1 allocs/op and 39.3 µs/op.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	doubled := Map(numbers, func(value int, _ int, _ []int) int {
//		return value * 2
//	})
//	// doubled == []int{2, 4, 6}
func Map[T any, V any](slice []T, fn MapFunc[T, V]) []V {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return make([]V, 0)
	}

	newList := make([]V, sliceLen)

	for i := 0; i < sliceLen; i++ {
		newList[i] = fn(slice[i], i, slice)
	}

	return newList
}

// MapStrict returns a new slice of the same element type containing the transformed values.
//
// Behavior:
//   - The original slice is not modified.
//   - Empty inputs return the original slice header.
//
// Performance Note:
//   - Current benchmark on this machine: 1 allocs/op and 36.8 µs/op.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	incremented := MapStrict(numbers, func(value int, _ int, _ []int) int {
//		return value + 1
//	})
//	// incremented == []int{2, 3, 4}
func MapStrict[T any](slice []T, fn MapFuncStrict[T]) []T {
	sliceLen := len(slice)
	if sliceLen == 0 {
		return slice
	}

	newList := make([]T, sliceLen)

	for i := 0; i < sliceLen; i++ {
		newList[i] = fn(slice[i], i, slice)
	}

	return newList
}

// Map returns a new array containing the transformed values of the receiver.
// It is the receiver-based form of Map.
func (a Array[T]) Map(fn MapFunc[T, any]) Array[any] {
	return Map(a, fn)
}

// MapStrict returns a new array of the same element type containing the transformed values of the receiver.
// It is the receiver-based form of MapStrict.
func (a Array[T]) MapStrict(fn MapFuncStrict[T]) Array[T] {
	return MapStrict(a, fn)
}
