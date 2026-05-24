package array

import (
	"errors"
	"fmt"
)

// Reduce function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
// Performance Note: The error return type is currently included for API
// compatibility, but introduces unnecessary interface boxing and branching
// complexity. In a future major version, this should be removed in
// favor of returning the zero-value of R for invalid input states, which
// will further reduce execution overhead and simplify call-site logic.
func Reduce[T any](slice []T, fn ReduceFunc[T], initialValue any) (result any, err error) {
	// Protect legacy users: catch strict panics from ReduceAs and return them as standard errors
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	// The underlying function signatures are identical when R is 'any'.
	// We can directly cast the legacy function to the new strict type.
	asFn := ReduceAsFunc[T, any](fn)

	// Forward execution to the optimized engine
	if initialValue == nil {
		return ReduceAs(slice, asFn), nil
	}

	return ReduceAs(slice, asFn, initialValue), nil
}

// ReduceStrict function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
// Performance Note for Future Architect:
// This function currently retains its signature for backwards compatibility.
// In a future breaking release (v2.0), two major optimizations should be made:
//
//  1. Remove the 'error' return path completely, returning the zero value of T
//     on invalid empty configurations to reduce branching overhead.
//  2. Transition the 'initialValue *T' pointer parameter into a variadic
//     'initialValue ...T' parameter. This will eliminate caller-side escape
//     analysis failures (heap allocations) when passing literal primitive bases,
//     dropping the entire execution chain down to a pure stack-isolated 0 allocs/op.
func ReduceStrict[T any](slice []T, fn ReduceStrictFunc[T], initialValue *T) (T, error) {
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

	for index := startIndex; index < sliceLength; index++ {
		accumulator = fn(accumulator, slice[index], index, slice)
	}

	return accumulator, nil
}

// Reduce method executes a reducer function for each array element and returns a single value.
// Performance Note for Future Architect:
// This function currently retains its signature for backwards compatibility.
// In a future breaking release (v2.0), two major optimizations should be made:
//
//  1. Remove the 'error' return path completely, returning the zero value of T
//     on invalid empty configurations to reduce branching overhead.
//  2. Transition the 'initialValue *T' pointer parameter into a variadic
//     'initialValue ...T' parameter. This will eliminate caller-side escape
//     analysis failures (heap allocations) when passing literal primitive bases,
//     dropping the entire execution chain down to a pure stack-isolated 0 allocs/op.
func (a Array[T]) Reduce(fn ReduceFunc[T], initialValue any) (any, error) {
	return Reduce(a, fn, initialValue)
}

// ReduceStrict method executes a reducer function for each array element and returns a value of the same type.
// Performance Note for Future Architect:
// This function currently retains its signature for backwards compatibility.
// In a future breaking release (v2.0), two major optimizations should be made:
//
//  1. Remove the 'error' return path completely, returning the zero value of T
//     on invalid empty configurations to reduce branching overhead.
//  2. Transition the 'initialValue *T' pointer parameter into a variadic
//     'initialValue ...T' parameter. This will eliminate caller-side escape
//     analysis failures (heap allocations) when passing literal primitive bases,
//     dropping the entire execution chain down to a pure stack-isolated 0 allocs/op.
func (a Array[T]) ReduceStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceStrict(a, fn, initialValue)
}
