package array

import "reflect"

// The Includes() function determines whether the array includes a certain element.
// NOTE: ❌ POTENTIAL PERFORMANCE TRAP
// Forces the caller to allocate an integer on the heap just to pass an index.
// Consider using variadic parameters better performance
func Includes[S ~[]T, T comparable](slice S, element T, start *int) bool {

	sliceLength := len(slice)
	if sliceLength == 0 {
		return false
	}

	startIndex := convertIndexSimply(sliceLength, *start)

	if startIndex == -1 {
		return false
	}

	for i := startIndex; i < sliceLength; i++ {
		if slice[i] == element {
			return true
		}
	}

	return false
}

// The Includes() method determines whether the array includes a certain element.
// Deprecated: Use ArrayComp.Includes or Includes function instead for better performence.
func (a Array[T]) Includes(element T, start *int) bool {
	sliceLength := len(a)
	startIndex := convertIndexSimply(sliceLength, *start)

	if startIndex == -1 {
		return false
	}

	for i := startIndex; i < sliceLength; i++ {
		value := a[i]

		if reflect.DeepEqual(element, value) {
			return true
		}
	}

	return false
}

// The Includes() method determines whether the array includes a certain element.
func (a ArrayComp[T]) Includes(element T, start ...int) bool {
	index := OptionalParam(start, 0)
	return Includes(a, element, &index)
}
