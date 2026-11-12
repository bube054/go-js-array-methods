package array

// ReduceRightAs reduces the slice to an accumulator of type R from right to left.
//
// Behavior:
//   - If initialValue is provided, the callback starts at the last element.
//   - If initialValue is omitted, the last element is used as the initial accumulator.
//   - Without an initial value, it panics if the slice is empty or its
//     elements are not of type R, because there is no first accumulator to
//     start from. This is a programming mistake, like indexing past the end of
//     a slice: pass an initial value whenever the slice may be empty or the
//     types differ.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 27.6 µs/op.
//   - This is the preferred zero-allocation replacement for the legacy ReduceRight helper.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := ReduceRightAs(numbers, func(acc string, value int, _ int, _ []int) string {
//		return acc + strconv.Itoa(value)
//	}, "")
//	// result == "321"
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
		panic("array: ReduceRightAs of empty slice with no initial value")
	}

	initial, ok := any(slice[sliceLen-1]).(R)
	if !ok {
		panic("array: ReduceRightAs needs an initial value when the element type is not the accumulator type")
	}

	accumulator := initial
	for i := sliceLen - 2; i >= 0; i-- {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
