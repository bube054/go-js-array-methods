package array

// The Reverse() function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original array.
func Reverse[T comparable](slice []T) []T {
	newList := []T{}

	for index := len(slice) - 1; index >= 0; index-- {
		item := slice[index]
		newList = append(newList, item)
	}

	return newList
}