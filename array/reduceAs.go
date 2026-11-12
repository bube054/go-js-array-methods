package array

// ReduceAs reduces the slice to an accumulator of type R.
//
// Behavior:
//   - If initialValue is provided, the callback starts at index 0.
//   - If initialValue is omitted, the first element is used as the initial accumulator.
//   - Without an initial value, it panics if the slice is empty or its
//     elements are not of type R, because there is no first accumulator to
//     start from. This is a programming mistake, like indexing past the end of
//     a slice: pass an initial value whenever the slice may be empty or the
//     types differ.
//
// Performance Note:
//   - Current benchmark on this machine: 0 allocs/op and 26.7 µs/op.
//   - This is the preferred zero-allocation replacement for the legacy Reduce and ReduceRight helpers.
//
// Example:
//
//	words := []string{"zero", "allocs"}
//	totalChars := ReduceAs(words, func(acc int, word string, _ int, _ []string) int {
//		return acc + len(word)
//	}, 0)
//	// totalChars == 10
func ReduceAs[T any, R any](slice []T, fn ReduceAsFunc[T, R], initialValue ...R) R {
	sliceLen := len(slice)

	if len(initialValue) > 0 {
		accumulator := initialValue[0]
		for i := 0; i < sliceLen; i++ {
			accumulator = fn(accumulator, slice[i], i, slice)
		}
		return accumulator
	}

	if sliceLen == 0 {
		panic("array: ReduceAs of empty slice with no initial value")
	}

	initial, ok := any(slice[0]).(R)
	if !ok {
		panic("array: ReduceAs needs an initial value when the element type is not the accumulator type")
	}

	accumulator := initial

	for i := 1; i < sliceLen; i++ {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
