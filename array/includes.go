package array

import "reflect"

// Includes function determines whether the array includes a certain element.
// NOTE: ❌ POTENTIAL PERFORMANCE TRAP
// Forces the caller to allocate an integer on the heap just to pass an index.
// Consider using variadic parameters better performance
func Includes[T comparable](slice []T, element T, start *int) bool {

	sliceLength := len(slice)
	if sliceLength == 0 {
		return false
	}

	startIndex := 0
	if start != nil {
		startIndex = convertIndexSimply(sliceLength, *start)
		if startIndex == -1 {
			return false
		}
	}

	for i := startIndex; i < sliceLength; i++ {
		if slice[i] == element {
			return true
		}
	}

	return false
}

// Includes method determines whether the array includes a certain element.
// Note: Use Includes function instead for better performence.
func (a Array[T]) Includes(element T, start *int) bool {
	sliceLength := len(a)
	startIndex := 0
	if start != nil {
		startIndex = convertIndexSimply(sliceLength, *start)
		if startIndex == -1 {
			return false
		}
	}

	// 3. Dual-Track Execution Engine
	if IsComparable(element) {
		for i := startIndex; i < sliceLength; i++ {
			if any(a[i]) == any(element) {
				return true
			}
		}
	} else {
		// SLOW PATH: Fallback for slices, maps, and uncomparable structs.
		for i := startIndex; i < sliceLength; i++ {
			if reflect.DeepEqual(element, a[i]) {
				return true
			}
		}
	}

	return false
}
