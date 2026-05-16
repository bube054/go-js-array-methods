package array

import (
	"fmt"
	"strconv"
	"strings"
)

// The Join() function returns an slice as a string. The Join() function does not change the original slice. Any separator can be specified. The default is comma (,).
func Join[S ~[]T, T any](slice S, separator *string) string {
	if len(slice) == 0 {
		return ""
	}

	sep := ","
	if separator != nil {
		sep = *separator
	}

	if strSlice, ok := any(slice).([]string); ok {
		return strings.Join(strSlice, sep)
	}

	// 3. The Builder Engine for generic/mixed types
	var b strings.Builder

	// Pre-allocate memory heuristic to prevent capacity resizing during the loop.
	// We guess roughly 8 bytes per element + space for the separators.
	b.Grow(len(slice)*8 + len(sep)*(len(slice)-1))

	// 4. Stream data directly into the builder buffer
	for i := 0; i < len(slice); i++ {
		if i > 0 {
			b.WriteString(sep)
		}

		// FAST PATHS: Direct hardware-level conversions completely bypassing reflect
		switch v := any(slice[i]).(type) {
		case string:
			b.WriteString(v)
		case int:
			b.WriteString(strconv.Itoa(v))
		case int64:
			b.WriteString(strconv.FormatInt(v, 10))
		case float64:
			b.WriteString(strconv.FormatFloat(v, 'g', -1, 64))
		case bool:
			b.WriteString(strconv.FormatBool(v))
		default:
			// SLOW PATH: Fallback to reflection only for complex structs or unhandled types
			b.WriteString(fmt.Sprint(v))
		}
	}

	return b.String()
}

// The Join() method joins all elements of the array into a single string using the provided separator.
// Elements are converted to their string form, so this works for any Array[T].
// If no separator is provided, a comma is used, matching JavaScript Array.join().
func (a Array[T]) Join(separator ...string) string {
	sep := OptionalParam(separator, ",")

	return Join(a, &sep)
}
