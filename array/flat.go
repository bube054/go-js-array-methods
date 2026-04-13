package array

// The Flat() function concatenates sub-arrays elements into a single new slice.
// The Flat() function returns a new slice containing all elements from the sub-arrays.
// The Flat() function does not change the original slice.
func Flat[T comparable](slice [][]T) []T {
	var flattenedSlice []T

	for _, subSlice := range slice {
		flattenedSlice = append(flattenedSlice, subSlice...)
	}

	return flattenedSlice
}
