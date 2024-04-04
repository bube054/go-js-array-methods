package array

// The ValueOf() function returns the slice itself. The ValueOf() function does not change the original slice. fruits. ValueOf() returns the same as fruits.
func ValueOf[T comparable](slice []T) []T {
	return slice
}