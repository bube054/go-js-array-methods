package array

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
