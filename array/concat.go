package array

// Concat function concatenates (joins) two or more slices.
// The concat() function returns a new slice, containing the joined slices.
// The concat() function does not change the existing slices.
func Concat[T any](slice1 []T, slices ...[]T) []T {
	totalLen := len(slice1)
	for i := 0; i < len(slices); i++ {
		totalLen += len(slices[i])
	}

	if len(slices) == 0 {
		return slice1
	}

	result := make([]T, totalLen)
	if len(slice1) > 0 {
		copy(result, slice1)
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

// Concat method concatenates the array with other slices and returns a new concatenated array.
func (a Array[T]) Concat(slices ...Array[T]) Array[T] {
	b := make([][]T, len(slices))
	for i, s := range slices {
		b[i] = s
	}
	return Array[T](Concat(a, b...))
}
