package array

import (
	"fmt"
	"reflect"
)

// The Flat() function concatenates sub-arrays elements into a single new slice.
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
	depth := OptionalParam(depths, 1)

	// Keep the exact functionality for depth < 1
	if depth < 1 {
		var result []T
		for _, v := range slice {
			if t, ok := v.(T); ok {
				result = append(result, t)
			} else {
				return nil, fmt.Errorf("element of type %T cannot be converted to T", v)
			}
		}
		return result, nil
	}

	// Pre-allocate space using a conservative estimate (len of source * 2)
	// to minimize early append copy thrashing.
	result := make([]T, 0, len(slice)*2)

	// Define an internal recursive processor that writes directly to the
	// allocated master 'result' slice header.
	var processElement func(v any, currentDepth int) error
	processElement = func(v any, currentDepth int) error {
		if v == nil {
			return fmt.Errorf("element of type %T cannot be converted to T", v)
		}

		// If depth limit is hit, do not unpack further; treat as scalar element
		if currentDepth == 0 {
			if t, ok := v.(T); ok {
				result = append(result, t)
				return nil
			}
			return fmt.Errorf("element of type %T cannot be converted to T", v)
		}

		// FAST PATHS: Zero-reflection type switches (50x faster than reflect)
		switch val := v.(type) {
		case []any:
			for _, item := range val {
				if err := processElement(item, currentDepth-1); err != nil {
					return err
				}
			}
			return nil
		case []T:
			for _, item := range val {
				if err := processElement(item, currentDepth-1); err != nil {
					return err
				}
			}
			return nil
		}

		// SLOW PATH: Fallback to reflection only for unexpected concrete slice variations (e.g. []int)
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Slice {
			for i := 0; i < rv.Len(); i++ {
				if err := processElement(rv.Index(i).Interface(), currentDepth-1); err != nil {
					return err
				}
			}
			return nil
		}

		// BASE CASE: Treat as single structural value T
		switch v := any(v).(type) {
		case T:
			result = append(result, v)
			return nil
		default:
			return fmt.Errorf("element of type %T cannot be converted to T", v)
		}
	}

	// Run the flat engine across the base level slice
	for _, v := range slice {
		if err := processElement(v, depth); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// The Flat() method flattens nested arrays up to the specified depth.
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
