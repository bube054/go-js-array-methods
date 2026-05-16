package array

// The Push() function adds new items to the end of an slice. The Push() function returns the new slice.
func Push[S ~[]T, T any](slice S, elements ...T) S {
	return append(slice, elements...)
}

// The Push() method adds one or more elements to the end of the array and returns the new array.
func (a Array[T]) Push(elements ...T) Array[T] {
	return Push(a, elements...)
}
