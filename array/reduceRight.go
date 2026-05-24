package array

import (
	"errors"
	"fmt"
)

// Reduce function executes a Reducer function for slice element. The ReduceRight() function works from right to left. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
func ReduceRight[T any](slice []T, fn ReduceFunc[T], initialValue any) (result any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	asFn := ReduceAsFunc[T, any](fn)

	if initialValue == nil {
		return ReduceRightAs(slice, asFn), nil
	}

	return ReduceRightAs(slice, asFn, initialValue), nil
}

// ReduceStrict function executes a Reducer function for slice element. The ReduceRightStrict() function works from right to left . The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
// Performance Note for Future Architect:
// This function currently retains its signature for backwards compatibility.
// In a future breaking release, two major optimizations should be made:
//
//  1. Remove the 'error' return path completely, returning the zero value of T
//     on invalid empty configurations to reduce branching overhead.
//  2. Transition the 'initialValue *T' pointer parameter into a variadic
//     'initialValue ...T' parameter. This will eliminate caller-side escape
//     analysis failures (heap allocations) when passing literal primitive bases,
//     dropping the entire execution chain down to a pure stack-isolated 0 allocs/op.
func ReduceRightStrict[T any](slice []T, fn ReduceStrictFunc[T], initialValue *T) (T, error) {
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
		accumulator = slice[sliceLength-1]
		startIndex = sliceLength - 2
	} else {
		accumulator = *initialValue
		startIndex = sliceLength - 1
	}

	for index := startIndex; index >= 0; index-- {
		accumulator = fn(accumulator, slice[index], index, slice)
	}

	return accumulator, nil
}

// The ReduceRight() method executes a reducer function for each array element (from right to left) and returns a single value.
// Performance Note: The error return type is currently included for API
// compatibility, but introduces unnecessary interface boxing and branching
// complexity. In a future major version, this should be removed in
// favor of returning the zero-value of R for invalid input states, which
// will further reduce execution overhead and simplify call-site logic.
func (a Array[T]) ReduceRight(fn ReduceFunc[T], initialValue any) (any, error) {
	return ReduceRight(a, fn, initialValue)
}

// ReduceRightStrict method executes a reducer function for each array element (from right to left) and returns a value of the same type.
// Performance Note: The error return type is currently included for API
// compatibility, but introduces unnecessary interface boxing and branching
// complexity. In a future major version, this should be removed in
// favor of returning the zero-value of R for invalid input states, which
// will further reduce execution overhead and simplify call-site logic.
func (a Array[T]) ReduceRightStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceRightStrict(a, fn, initialValue)
}
