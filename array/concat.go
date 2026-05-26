package array

// Concat returns a new slice containing the elements of slice1 and the provided suffix slices.
//
// Behavior:
//   - The input slices are not modified.
//   - The returned slice is a new backing array containing the concatenated values.
//
// Performance Note:
//   - Current benchmark on this machine: 1 allocs/op and 31.7 µs/op.
//
// Example:
//
//	left := []int{1, 2}
//	right := []int{3, 4}
//	result := Concat(left, right)
//	// result == []int{1, 2, 3, 4}
func Concat[T any](slice1 []T, slices ...[]T) []T {
	totalLen := len(slice1)
	for i := 0; i < len(slices); i++ {
		totalLen += len(slices[i])
	}

	result := make([]T, totalLen)

	copy(result, slice1)

	if len(slices) == 0 {
		return result
	}

	pos := len(slice1)
	for i := 0; i < len(slices); i++ {
		s := slices[i]
		if len(s) == 0 {
			continue
		}
		copy(result[pos:], s)
		pos += len(s)
	}

	return result
}

// Concat returns a new array containing the concatenated values of the receiver and the provided arrays.
// It is the receiver-based form of Concat.
func (a Array[T]) Concat(slices ...Array[T]) Array[T] {
	b := make([][]T, len(slices))
	for i, s := range slices {
		b[i] = s
	}
	return Array[T](Concat(a, b...))
}
