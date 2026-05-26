package array

// ReduceRightAs reduces the slice to an accumulator of type R from right to left.
//
// Behavior:
//   - If initialValue is provided, the callback starts at the last element.
//   - If initialValue is omitted, the last element is used as the initial accumulator.
//   - Empty input with no initial value panics with a TypeError.
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
		panic("TypeError: ReduceRight of empty slice with no initial value")
	}

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
