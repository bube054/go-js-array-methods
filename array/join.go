package array

import "strings"

// The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,). Join function only works for strings.
func Join(slice []string, separator *string) string {
	if separator == nil {
		defaultSeparator := ","
		separator = &defaultSeparator
	}

	return strings.Join(slice, *separator)
}
