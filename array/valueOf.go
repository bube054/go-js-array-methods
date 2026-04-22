package array

// The ValueOf() function returns the slice itself. The ValueOf() function does not change the original slice. fruits. ValueOf() returns the same as fruits.
func ValueOf[T comparable](slice []T) []T {
	return slice
}

// The ValueOf() method returns the array itself (identity function).
func (a Array[T]) ValueOf() Array[T] {
	return Array[T](ValueOf([]T(a)))
}
