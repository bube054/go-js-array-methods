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

	if depth < 1 {
		// When depth < 1, convert direct elements to []T without further flattening
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

	var result []T
	for _, v := range slice {
		if v != nil {
			rv := reflect.ValueOf(v)
			if rv.Kind() == reflect.Slice {
				sliceAsAny := make([]any, rv.Len())
				for i := 0; i < rv.Len(); i++ {
					sliceAsAny[i] = rv.Index(i).Interface()
				}
				flatSlice, err := Flat[T](sliceAsAny, depth-1)
				if err != nil {
					return nil, err
				}
				result = append(result, flatSlice...)
				continue
			}
		}

		switch v := any(v).(type) {
		case T:
			result = append(result, v)
		default:
			return nil, fmt.Errorf("element of type %T cannot be converted to T", v)
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
