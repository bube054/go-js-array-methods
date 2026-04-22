package array

// The With() function updates a specified slice element. The With() function returns a new slice. The With() function does not change the original slice.
func With[T comparable](slice []T, index int, value T) ([]T, error) {
	newIndex, err := ConvertIndex(slice, index, "index")

	if err != nil {
		return slice, err
	}

	sliceLength := len(slice)
	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newSlice[newIndex] = value

	return newSlice, nil
}

// The With() method returns a new array with the element at the given index replaced with the given value.
func (a Array[T]) With(index int, value T) (Array[T], error) {
	result, err := With([]T(a), index, value)
	return Array[T](result), err
}
