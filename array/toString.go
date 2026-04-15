package array

// The ToString() function returns a string with slice values separated by commas. The ToString() function does not change the original slice.
func ToString[T](slice []T) string {
	return Join(slice, nil)
}
