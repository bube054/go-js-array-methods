package array

// The ToString() function returns a string with slice values separated by commas. The ToString() function does not change the original slice. The ToString function only works for a slice of strings.
func ToString(slice []string) string {
	return Join(slice, nil)
}