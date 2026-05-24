package array

import (
	"fmt"
)

// Flat function concatenates sub-arrays elements into a single new slice.
// The Flat() function returns a new slice containing all elements from the sub-arrays.
// The Flat() function does not change the original slice. The depth parameter specifies how deep a nested array structure should be flattened. The default is 1.
// The Flat() function can handle slices of any type, but all elements must be of the same type T. If an element cannot be converted to type T, an error is returned.
// Example usage:
//
//	input := []any{1, []int{2, 3}, 4}
//	result, err := Flat[int](input)
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Println(result) // Output: [1 2 3 4]
//	}
func Flat[T any](slice []any, depths ...int) ([]T, error) {
	depth := optionalParam(depths, 1)
	sliceLen := len(slice)

	// Keep the exact functionality for depth < 1
	if depth < 1 {
		var result []T
		for i := 0; i < sliceLen; i++ {
			if t, ok := slice[i].(T); ok {
				result = append(result, t)
			} else {
				return nil, fmt.Errorf("element of type %T cannot be converted to T", slice[i])
			}
		}
		return result, nil
	}

	if sliceLen == 0 {
		var result []T
		return result, nil
	}

	// Pre-allocate space using your conservative estimate
	result := make([]T, 0, sliceLen*2)

	// Pass a pointer to result (*[]T) to ensure zero closure allocations
	if err := flatProcessor(&result, slice, depth); err != nil {
		return nil, err
	}

	return result, nil
}

// Flat method flattens nested arrays up to the specified depth.
// The Flat() method returns a new array with the sub-array elements concatenated into it. The depth parameter specifies how deep a nested array structure should be flattened. The default is 1.
// Until Generic Methods are implemented in Go, the Flat() method can only return Array[any] since it needs to handle slices of any type. Once Generic Methods are available, we can update the Flat() method to return Array[T] and handle type assertions accordingly.
func (a Array[T]) Flat(depths ...int) (Array[any], error) {
	// Convert Array[T] to []any for flattening
	anySlice := make([]any, len(a))
	for i, v := range a {
		anySlice[i] = v
	}
	result, err := Flat[any](anySlice, depths...)
	return Array[any](result), err
}
