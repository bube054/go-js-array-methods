package array

import (
	"errors"
	"fmt"
)

// Reduce is deprecated.
//
// Roadmap Notice:
//   - This legacy any-backed signature will be removed in the next major release.
//   - Prefer ReduceAs for zero-allocation, typed accumulator pipelines.
//
// Performance Note:
//   - Current benchmark on this machine: 9,977 allocs/op and 271.7 µs/op.
//   - The boxed any accumulator is not suitable for hot loops.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	total, err := Reduce(numbers, func(acc any, value int, _ int, _ []int) any {
//		return acc.(int) + value
//	}, 0)
//	// total == 10
//	// err == nil
func Reduce[T any](slice []T, fn ReduceFunc[T], initialValue any) (result any, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	asFn := ReduceAsFunc[T, any](fn)

	if initialValue == nil {
		return ReduceAs(slice, asFn), nil
	}

	return ReduceAs(slice, asFn, initialValue), nil
}

// ReduceStrict reduces the slice to a value of the same type.
//
// Behavior:
//   - If initialValue is nil, the first element becomes the accumulator and the loop starts at index 1.
//   - If initialValue is provided, it is used as the accumulator and the loop starts at index 0.
//   - Empty input with no initial value returns an error.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 27.8 µs/op.
//   - Roadmap Notice: the pointer-based initialValue parameter is retained for compatibility; a future major release may switch to variadic initialValue ...T.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	initial := 0
//	total, err := ReduceStrict(numbers, func(acc int, value int, _ int, _ []int) int {
//		return acc + value
//	}, &initial)
//	// total == 6
//	// err == nil
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

// Reduce is deprecated.
// It is the receiver-based form of Reduce.
func (a Array[T]) Reduce(fn ReduceFunc[T], initialValue any) (any, error) {
	return Reduce(a, fn, initialValue)
}

// ReduceStrict reduces the receiver to a value of the same type.
// It is the receiver-based form of ReduceStrict.
func (a Array[T]) ReduceStrict(fn ReduceStrictFunc[T], initialValue *T) (T, error) {
	return ReduceStrict(a, fn, initialValue)
}
