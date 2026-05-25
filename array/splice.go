package array

import "fmt"

// Splice returns a new slice with elements inserted, removed, or replaced.
//
// Behavior:
//   - The original slice is not modified.
//   - The optional howMany pointer controls how many elements are removed; nil removes to the end.
//   - Invalid indexes return an error.
//
// Performance Note:
//   - The implementation builds the result in a single allocation and copies each segment once.
//   - Roadmap Notice: the howMany pointer is retained for compatibility; a future major release should switch to a plain int to avoid avoidable escape-analysis pressure.
//
// Example:
//
//	numbers := []int{1, 2, 3, 4}
//	count := 1
//	result, err := Splice(numbers, 1, &count, 9)
//	// result == []int{1, 9, 3, 4}
//	// err == nil
func Splice[T any](slice []T, index int, howMany *int, elements ...T) ([]T, error) {
	sliceLen := len(slice)

	newIndex := convertIndex(sliceLen, index)
	if newIndex == -1 {
		return slice, fmt.Errorf("Invalid index %d", index)
	}

	count := 0
	if howMany != nil {
		count = *howMany
		if count < 0 {
			count = 0
		}
	} else {
		count = sliceLen - newIndex
	}

	if newIndex+count > sliceLen {
		count = sliceLen - newIndex
	}

	insertLen := len(elements)
	tailIdx := newIndex + count
	tailLen := sliceLen - tailIdx
	finalLen := newIndex + insertLen + tailLen

	if finalLen == 0 {
		return make([]T, 0), nil
	}

	result := make([]T, finalLen)
	copy(result[0:newIndex], slice[0:newIndex])
	if insertLen > 0 {
		copy(result[newIndex:newIndex+insertLen], elements)
	}
	if tailLen > 0 {
		copy(result[newIndex+insertLen:], slice[tailIdx:])
	}

	return result, nil
}

// Splice returns a new array with elements inserted, removed, or replaced.
// It is the receiver-based form of Splice.
func (a Array[T]) Splice(index int, howMany *int, items ...T) (Array[T], error) {
	spliced, err := Splice(a, index, howMany, items...)
	return Array[T](spliced), err
}
