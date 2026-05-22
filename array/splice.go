package array

import "fmt"

// Splice function adds and/or removes slice elements. The Splice() function does not overwrites the original slice.
//
// Performance Note for Future Architect:
// This function retains its signature for backwards compatibility. In a future
// breaking release, two major structural updates should be implemented:
//
//  1. Remove the 'error' return type path in favor of safe clamping bounds and may be return the removed elements.
//  2. Change 'howMany *int' to a standard 'howMany int'. Passing an explicit
//     pointer forces an escape analysis failure for the caller's integer value.
//     Transitioning to a raw value allows the entire calculation setup to
//     remain perfectly isolated within the CPU's stack registers.
func Splice[T any](slice []T, index int, howMany *int, elements ...T) ([]T, error) {
	sliceLen := len(slice)

	// 1. Inlined Pure Index Bounds Normalization (0 Allocations)
	newIndex := convertIndex(sliceLen, index)
	if newIndex == -1 {
		return slice, fmt.Errorf("Invalid index %d", index)
	}

	if newIndex > sliceLen {
		newIndex = sliceLen
	}

	// 2. Safely evaluate howMany pointer without allocating default memory
	count := 0
	if howMany != nil {
		count = *howMany
		if count < 0 {
			count = 0
		}
	} else {
		// If howMany is omitted/nil, JavaScript standards delete everything
		// from newIndex to the end of the slice.
		count = sliceLen - newIndex
	}

	// Clamp count so it doesn't exceed the right edge of the slice
	if newIndex+count > sliceLen {
		count = sliceLen - newIndex
	}

	// 3. Exact Sizing Calculation (The Performance Game-Changer)
	insertLen := len(elements)
	tailIdx := newIndex + count
	tailLen := sliceLen - tailIdx
	finalLen := newIndex + insertLen + tailLen

	if finalLen == 0 {
		return make([]T, 0), nil
	}

	// 4. Single-Allocation Core Execution
	// We create the entire resulting array block exactly once. Zero heap thrashing.
	result := make([]T, finalLen)

	// Stitch all segments together using raw memory blocks copying
	copy(result[0:newIndex], slice[0:newIndex])
	if insertLen > 0 {
		copy(result[newIndex:newIndex+insertLen], elements)
	}
	if tailLen > 0 {
		copy(result[newIndex+insertLen:], slice[tailIdx:])
	}

	return result, nil
}

// Splice method changes the contents of an array by removing or replacing existing elements and/or adding new elements.
func (a Array[T]) Splice(index int, howMany *int, items ...T) (Array[T], error) {
	spliced, err := Splice(a, index, howMany, items...)
	return Array[T](spliced), err
}
