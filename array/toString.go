package array

// The ToString() function returns a string with slice values separated by commas. The ToString() function does not change the original slice.
func ToString[T any](slice []T) string {
	return Join(slice, nil)
}

// The ToString() method returns a string representation of the array.
// It mirrors JavaScript Array.toString() by joining elements with commas.
func (a Array[T]) ToString() string {
	return a.Join()
}
