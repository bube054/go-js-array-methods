package array

func ReduceRightAs[T any, R any](slice []T, fn ReduceAsFunc[T, R], initialValue ...R) R {
	sliceLen := len(slice)

	var accumulator R
	hasInitial := false
	if len(initialValue) > 0 {
		accumulator = initialValue[0]
		hasInitial = true
	}

	if sliceLen == 0 {
		return accumulator
	}

	startIndex := sliceLen - 1
	if !hasInitial {
		if initial, ok := any(slice[sliceLen-1]).(R); ok {
			accumulator = initial
			startIndex = sliceLen - 2
		} else {
			return accumulator
		}
	}

	for i := startIndex; i >= 0; i-- {
		accumulator = fn(accumulator, slice[i], i, slice)
	}

	return accumulator
}
