package array

// The concat() function concatenates (joins) two or more slices.
// The concat() function returns a new slice, containing the joined slices.
// The concat() function does not change the existing slices.
func Concat[T any](slice1 []T, slices ...[]T) []T {
	totalLen := len(slice1)
	for _, s := range slices {
		totalLen += len(s)
	}

	result := make([]T, 0, totalLen)

	result = append(result, slice1...)

	for _, s := range slices {
		result = append(result, s...)
	}

	return result
}

// The Concat() method concatenates the array with other slices and returns a new concatenated array.
func (a Array[T]) Concat(slices ...Array[T]) Array[T] {
	b := make([][]T, len(slices))
	for i, s := range slices {
		b[i] = []T(s)
	}
	return Array[T](Concat([]T(a), b...))
}
