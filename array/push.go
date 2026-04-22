package array

// The Push() function adds new items to the end of an slice. The Push() function returns the new slice.
func Push[T comparable](slice []T, elements ...T) []T {
	return append(slice, elements...)
}

// The Push() method adds one or more elements to the end of the array and returns the new array.
func (a Array[T]) Push(elements ...T) Array[T] {
	return Array[T](Push([]T(a), elements...))
}
