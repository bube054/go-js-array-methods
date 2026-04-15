package array

import (
	"fmt"
	"strings"
)

// The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings.
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
