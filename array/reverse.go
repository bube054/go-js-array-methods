package array

// Reverse function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original array.
func Reverse[T any](slice []T) []T {
	newList := make([]T, 0, len(slice))

	for index := len(slice) - 1; index >= 0; index-- {
		item := slice[index]
		newList = append(newList, item)
	}

	return newList
}

// Reverse method reverses the array and returns a new reversed array.
func (a Array[T]) Reverse() Array[T] {
	return Array[T](Reverse(a))
}
