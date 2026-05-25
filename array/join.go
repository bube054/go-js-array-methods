package array

import (
	"strconv"
	"strings"
)

// Join returns the slice as a string using the provided separator.
//
// Behavior:
//   - The original slice is not modified.
//   - When separator is nil, a comma is used.
//
// Performance Note:
//   - Current benchmark on this machine: 1645 allocs/op and 66.0 µs/op.
//   - Common primitive types are formatted through direct string conversions; complex values fall back to fmt.Sprint.
//
// Example:
//
//	numbers := []int{1, 2, 3}
//	result := Join(numbers, nil)
//	// result == "1,2,3"
func Join[T any](slice []T, separator *string) string {
	if len(slice) == 0 {
		return ""
	}

	sep := ","
	if separator != nil {
		sep = *separator
	}

	switch s := any(slice).(type) {
	case []string:
		return strings.Join(s, sep)

	case []int:
		var b strings.Builder
		b.Grow(len(s)*4 + len(sep)*(len(s)-1))
		var buf [32]byte 

		b.Write(strconv.AppendInt(buf[:0], int64(s[0]), 10))
		for _, v := range s[1:] {
			b.WriteString(sep)
			b.Write(strconv.AppendInt(buf[:0], int64(v), 10))
		}
		return b.String()

	case []int64:
		var b strings.Builder
		b.Grow(len(s)*4 + len(sep)*(len(s)-1))
		var buf [32]byte

		b.Write(strconv.AppendInt(buf[:0], s[0], 10))
		for _, v := range s[1:] {
			b.WriteString(sep)
			b.Write(strconv.AppendInt(buf[:0], v, 10))
		}
		return b.String()

	case []float64:
		var b strings.Builder
		b.Grow(len(s)*8 + len(sep)*(len(s)-1))
		var buf [64]byte

		b.Write(strconv.AppendFloat(buf[:0], s[0], 'g', -1, 64))
		for _, v := range s[1:] {
			b.WriteString(sep)
			b.Write(strconv.AppendFloat(buf[:0], v, 'g', -1, 64))
		}
		return b.String()

	case []bool:
		var b strings.Builder
		b.Grow(len(s)*4 + len(sep)*(len(s)-1))

		b.WriteString(strconv.FormatBool(s[0]))
		for _, v := range s[1:] {
			b.WriteString(sep)
			b.WriteString(strconv.FormatBool(v))
		}
		return b.String()
	}
	// 2. Fallback Path (For custom structs, `[]any`, or unhandled types)
	var b strings.Builder
	b.Grow(len(slice)*8 + len(sep)*(len(slice)-1))

	// Process the first element outside the loop to eliminate `if i > 0` checks
	writeAny(&b, slice[0])
	for i := 1; i < len(slice); i++ {
		b.WriteString(sep)
		writeAny(&b, slice[i])
	}

	return b.String()
}

// Join returns a string representation of the receiver using the provided separator.
// It is the receiver-based form of Join.
func (a Array[T]) Join(separator ...string) string {
	sep := optionalParam(separator, ",")

	return Join(a, &sep)
}
