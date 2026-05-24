package array

// ReduceRightAs executes a user-provided callback function on each element of the slice,
// from right to left (backward), passing in the return value from the calculation on the succeeding element.
// It uses dual-generic type parameters [T, R] to reduce a slice of type T backward into a completely
// distinct accumulator type R, achieving maximum performance with zero runtime memory allocations.
//
// If an initial value is provided, the accumulator is initialized to that value, and the reduction
// starts at the trailing element (len-1). If initialValue is omitted, the last element of the slice is used
// as the initial accumulator value, and the loop starts at the second-to-last element (len-2).
//
// Behavior:
//   - If the slice is empty and an initial value is provided, it returns the initial value.
//   - If the slice is empty and no initial value is provided, it panics with a TypeError.
//   - If no initial value is provided and the slice element type T cannot be structurally verified
//     as the accumulator type R, it panics with a TypeError.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	csv := array.ReduceRightAs(numbers, func(acc string, num int, i int, s []int) string {
//		return acc + strconv.Itoa(num)
//	}, "")
//	// csv = "321"
func ReduceRightAs[T any, R any](slice []T, fn ReduceAsFunc[T, R], initialValue ...R) R {
	sliceLen := len(slice)

	if len(initialValue) > 0 {
		accumulator := initialValue[0]
		for i := sliceLen - 1; i >= 0; i-- {
			accumulator = fn(accumulator, slice[i], i, slice)
		}
		return accumulator
	}

	if sliceLen == 0 {
		panic("TypeError: ReduceRight of empty slice with no initial value")
	}

	// Runtime assertion: verify the trailing element can safely act as accumulator type R.
	// This invokes a micro-allocation ONLY when the initial value is omitted.
	initial, ok := any(slice[sliceLen-1]).(R)
	if !ok {
		panic("TypeError: slice element type does not match accumulator type and no initial value provided")
	}

	accumulator := initial
	for i := sliceLen - 2; i >= 0; i-- {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
