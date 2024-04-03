package array

// The Push() function adds new items to the end of an slice. The Push() function returns the new slice.
func Push[T comparable](s []T, e... T) []T {
	return append(s, e...)
}
