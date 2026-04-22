package array

import (
	"fmt"
	"strings"
)

// The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,).
func Join[T any](slice []T, separator *string) string {
	if separator == nil {
		defaultSeparator := ","
		separator = &defaultSeparator
	}

	strSlice := make([]string, len(slice))
	for i, v := range slice {
		strSlice[i] = fmt.Sprint(v)
	}

	return strings.Join(strSlice, *separator)
}

// The Join() method joins all elements of the array into a single string using the provided separator.
// Elements are converted to their string form, so this works for any Array[T].
// If no separator is provided, a comma is used, matching JavaScript Array.join().
func (a Array[T]) Join(separator ...string) string {
	sep := OptionalParam(separator, ",")

	return Join([]T(a), &sep)
}
