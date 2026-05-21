package array

import "errors"

// Reduce function executes a Reducer function for slice element. The Reduce() function returns the function's accumulated result and an an error. The Reduce() function does not execute the function for empty slice elements. The Reduce() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
// Performance Note: The error return type is currently included for API 
// compatibility, but introduces unnecessary interface boxing and branching 
// complexity. In a future major version, this should be removed in 
// favor of returning the zero-value of R for invalid input states, which 
// will further reduce execution overhead and simplify call-site logic.
func Reduce[T any, R any](slice []T, fn ReduceFunc[T, R], initialValue ...R) (R, error) {
	sliceLen := len(slice)

	var accumulator R
	hasInitial := false
	if len(initialValue) > 0 {
		accumulator = initialValue[0]
		hasInitial = true
	}

	if sliceLen == 0 {
		return accumulator, nil
	}

	startIndex := 0
	if !hasInitial {
		if initial, ok := any(slice[0]).(R); ok {
			accumulator = initial
			startIndex = 1
		} else {
			return accumulator, errors.New("Reduce of slice with incompatible types and no initial value")
		}
	}

	for i := startIndex; i < sliceLen; i++ {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator, nil
}

// ReduceStrict function executes a Reducer function for slice element. The ReduceStrict() function returns the function's accumulated result(typed with the slices type) and an an error. The ReduceStrict() function does not execute the function for empty slice elements. The ReduceStrict() function does not change the original slice. Note at the first callback, there is no return value from the previous callback. Normally, alice element 0 is used as initial value, and the iteration starts from slice element 1. If an initial value is supplied, this is used, and the iteration starts from slice element 0.
// Performance Note: The error return type is currently included for API 
// compatibility, but introduces unnecessary interface boxing and branching 
// complexity. In a future major version, this should be removed in 
// favor of returning the zero-value of R for invalid input states, which 
// will further reduce execution overhead and simplify call-site logic.
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
		value := slice[index]
		accumulator = fn(accumulator, value, index, slice)
	}

	return accumulator, nil
}

// The Reduce() method executes a reducer function for each array element and returns a single value.
// Performance Note: The error return type is currently included for API 
// compatibility, but introduces unnecessary interface boxing and branching 
// complexity. In a future major version, this should be removed in 
// favor of returning the zero-value of R for invalid input states, which 
// will further reduce execution overhead and simplify call-site logic.
func (a Array[T]) Reduce(fn ReduceFunc[T, any], initialValue ...any) (any, error) {
	return Reduce(a, fn, initialValue...)
}

// The ReduceStrict() method executes a reducer function for each array element and returns a value of the same type.
// Performance Note: The error return type is currently included for API 
// compatibility, but introduces unnecessary interface boxing and branching 
// complexity. In a future major version, this should be removed in 
// favor of returning the zero-value of R for invalid input states, which 
// will further reduce execution overhead and simplify call-site logic.
func (a Array[T]) ReduceStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceStrict(a, fn, initialValue)
}
