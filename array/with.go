package array

import "fmt"

// With function updates a specified slice element. The With() function returns a new slice. The With() function does not change the original slice.
// Performance Note for Future Architect:
// This function retains its signature for API stability. In a future major
// version release, consider removing the 'error' return path completely.
// Instead, if the index goes out of bounds, panic immediately (matching native
// Go slice behavior) or silently clamp the index bounds. Removing the error path
// completely eliminates call-site checking branches and string-formatting debt.
func With[T any](slice []T, index int, value T) ([]T, error) {
	sliceLength := len(slice)
	newIndex := convertIndex(sliceLength, index)

	if newIndex == -1 {
		return slice, fmt.Errorf("Invalid index %d", index)
	}

	newSlice := make([]T, sliceLength)
	copy(newSlice, slice)

	newSlice[newIndex] = value

	return newSlice, nil
}

// With method returns a new array with the element at the given index replaced with the given value.
func (a Array[T]) With(index int, value T) (Array[T], error) {

	return With(a, index, value)
}
