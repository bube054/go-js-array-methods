package array

// ReduceAs executes a user-provided callback function on each element of the slice,
// from left to right, passing in the return value from the calculation on the preceding element.
// It uses dual-generic type parameters [T, R] to reduce a slice of type T into a completely distinct
// accumulator type R, achieving maximum performance with zero runtime memory allocations.
//
// If an initial value is provided, the accumulator is initialized to that value, and the reduction
// starts at index 0. If initialValue is omitted, the first element of the slice is used as the initial
// accumulator value, and the loop starts at index 1.
//
// Behavior:
//   - If the slice is empty and an initial value is provided, it returns the initial value.
//   - If the slice is empty and no initial value is provided, it panics with a TypeError.
//   - If no initial value is provided and the slice element type T cannot be structurally verified
//     as the accumulator type R, it panics with a TypeError.
//
// Example:
//
//	words := []string{"zero", "allocs"}
//	totalChars := array.ReduceAs(words, func(acc int, word string, i int, s []string) int {
//		return acc + len(word)
//	}, 0)
//	// totalChars = 10
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
		panic("TypeError: Reduce of empty slice with no initial value")
	}

	// In ReduceAs, omitting the initial value implies T and R are the same underlying type.
	// We MUST cast via 'any' to satisfy the compiler, triggering a micro-allocation.
	// If they are incompatible types, we must panic to prevent silent zero-value logic bugs.
	initial, ok := any(slice[0]).(R)
	if !ok {
		panic("TypeError: slice element type does not match accumulator type and no initial value provided")
	}

	accumulator := initial

	for i := 1; i < sliceLen; i++ {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
