package array

// The concat() function concatenates (joins) two or more slices.
// The concat() function returns a new slice, containing the joined slices.
// The concat() function does not change the existing slices.
func Concat[T comparable](slice1 []T, slice2 ...[]T) []T {
	var all []T = slice1

	for _, list := range slice2 {
		all = append(all, list...)
	}

	return all
}

// The Concat() method concatenates the array with other slices and returns a new concatenated array.
func (a Array[T]) Concat(slices ...Array[T]) Array[T] {
	b := make([][]T, len(slices))
	for i, s := range slices {
		b[i] = []T(s)
	}
	return Array[T](Concat([]T(a), b...))
}
