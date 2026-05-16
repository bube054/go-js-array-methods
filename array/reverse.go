package array

// The Reverse() function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original array.
func Reverse[S ~[]T, T any](slice S) S {
	newList := make(S, 0, len(slice))

	for index := len(slice) - 1; index >= 0; index-- {
		item := slice[index]
		newList = append(newList, item)
	}

	return newList
}

// The Reverse() method reverses the array and returns a new reversed array.
func (a Array[T]) Reverse() Array[T] {
	return Reverse(a)
}
