package array

func ReduceAs[T any, R any](slice []T, fn ReduceAsFunc[T, R], initialValue ...R) R {
	sliceLen := len(slice)

	var accumulator R
	hasInitial := false
	if len(initialValue) > 0 {
		accumulator = initialValue[0]
		hasInitial = true
	}

	if sliceLen == 0 {
		if !hasInitial {
			var zero R
			return zero
		}
		return accumulator
	}

	startIndex := 0
	if !hasInitial {
		if initial, ok := any(slice[0]).(R); ok {
			accumulator = initial
			startIndex = 1
		} else {
			return accumulator
		}
	}

	for i := startIndex; i < sliceLen; i++ {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
