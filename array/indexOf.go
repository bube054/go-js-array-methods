package array

import "reflect"

func IndexOf[T comparable](s []T, e T) int {
	for index, val := range s {
		if reflect.DeepEqual(val, e) {
			return index
		}
	}

	return -1
}
