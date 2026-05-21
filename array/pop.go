package array

// Pop function removes (pops) the last element of an slice. The Pop() function does not change the original slice. The Pop() function returns the new slice without the last element and a pointer of removed element.
// Note on Architecture & Performance:
// Pop currently returns a pointer to the removed element (*T), which forces the Go compiler to move the element to the heap via escape analysis. To achieve a zero-allocation pipeline and prevent underlying array retention, future API iterations should return the element directly by value T instead of a pointer.
func Pop[T any](slice []T) ([]T, *T) {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return slice, nil
	}

	lastElement := slice[sliceLength-1]

	return slice[:sliceLength-1], &lastElement
}

// Pop method removes the last element from the array and returns the modified array and the removed element.
func (a Array[T]) Pop() (Array[T], *T) {

	return Pop(a)
}
