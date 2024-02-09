package array

func Includes[T comparable](s []T, e T) bool {
	index := IndexOf[T](s, e)

	if index < 0 {
		return false
	} else {
		return true
	}
}
