package array

import (
	"errors"
	"fmt"
)

// ReduceRight is deprecated.
//
// Roadmap Notice:
//   - This legacy any-backed signature will be removed in the next major release.
//   - Prefer ReduceRightAs for zero-allocation, typed accumulator pipelines.
//
// Performance Note:
//   - This legacy wrapper boxes values through any and should be avoided in hot loops.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	total, err := ReduceRight(numbers, func(acc any, value int, _ int, _ []int) any {
//		return acc.(int) + value
//	}, 0)
//	// total == 6
//	// err == nil
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

// ReduceRightStrict reduces the slice to a value of the same type from right to left.
//
// Behavior:
//   - If initialValue is nil, the last element becomes the accumulator and the loop starts at index len(slice)-2.
//   - If initialValue is provided, it is used as the accumulator and the loop starts at the last element.
//   - Empty input with no initial value returns an error.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 27.6 µs/op when routed through ReduceRightAs.
//   - Roadmap Notice: the pointer-based initialValue parameter is retained for compatibility; a future major release may switch to variadic initialValue ...T.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	initial := 0
//	total, err := ReduceRightStrict(numbers, func(acc int, value int, _ int, _ []int) int {
//		return acc + value
//	}, &initial)
//	// total == 6
//	// err == nil
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

// ReduceRight is deprecated.
// It is the receiver-based form of ReduceRight.
func (a Array[T]) ReduceRight(fn ReduceFunc[T], initialValue any) (any, error) {
	return ReduceRight(a, fn, initialValue)
}

// ReduceRightStrict reduces the receiver to a value of the same type from right to left.
// It is the receiver-based form of ReduceRightStrict.
func (a Array[T]) ReduceRightStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceRightStrict(a, fn, initialValue)
}
