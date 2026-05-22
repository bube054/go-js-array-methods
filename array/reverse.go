package array

// Reverse function Reverses the order of the elements in an slice. The Reverse() function does not overwrites the original array.
func Reverse[T any](slice []T) []T {
	sliceLen := len(slice)
	var newList []T
	if sliceLen == 0 {
		return newList
	}

	newList = make([]T, sliceLen)

	for i, j := sliceLen-1, 0; i >= 0; i, j = i-1, j+1 {
		newList[j] = slice[i]
	}

	return newList
}

// Reverse method reverses the array and returns a new reversed array.
func (a Array[T]) Reverse() Array[T] {
	return Array[T](Reverse(a))
}
