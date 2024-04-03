package array

import "errors"

// The Reduce() function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
func ReduceRight[T comparable](slice []T, fn ReduceFunc[T], initialValue any) (any, error) {
	var sliceLength = len(slice)

	if sliceLength == 0 && initialValue == nil {
		return nil, errors.New("Reduce of empty slice with no initial value")
	}

	if sliceLength == 0 && initialValue != nil {
		return initialValue, nil
	}

	var accumulator any
	var startIndex int

	if initialValue == nil {
		accumulator = slice[0]
		startIndex = 1
	} else {
		accumulator = initialValue
		startIndex = 0
	}

	for index := sliceLength - 1; index >= startIndex; index-- {
		value := slice[index]
		accumulator = fn(accumulator, value, index, slice)
	}

	return accumulator, nil
}

// The ReduceStrict() function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
func ReduceRightStrict[T comparable](slice []T, fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	var sliceLength = len(slice)

	if sliceLength == 0 && initialValue == nil {
		var defaultVal T
		return defaultVal, errors.New("Reduce of empty slice with no initial value")
	}

	if sliceLength == 0 && initialValue != nil {
		return *initialValue, nil
	}

	var accumulator T
	var startIndex int

	if initialValue == nil {
		accumulator = slice[0]
		startIndex = 1
	} else {
		accumulator = *initialValue
		startIndex = 0
	}

	for index := sliceLength - 1; index >= startIndex; index-- {
		value := slice[index]
		accumulator = fn(accumulator, value, index, slice)
	}

	return accumulator, nil
}
