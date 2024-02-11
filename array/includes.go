package array

// Includes checks if a given element is present in the array.
func Includes[T comparable](s []T, e T) bool {
	index := IndexOf[T](s, e)

	if index < 0 {
		return false
	} else {
		return true
	}
}
