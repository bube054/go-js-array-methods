package array

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// optionalParam is a helper function to handle optional parameters with default values.
func optionalParam[T any](params []T, defaultValue T) T {
	if len(params) > 0 {
		return params[0]
	}
	return defaultValue
}

// OptionalParam is an exported version of optionalParam for use in other packages if needed.
// Deprecated: This function is not intended for public use and may be removed in future versions. Use local implementations within package to avoid interface allocations.
func OptionalParam[T any](params []T, defaultValue T) T {
	return optionalParam(params, defaultValue)
}

// ConvertIndex converts negative JS-style indices to absolute Go slice indices.
// It achieves zero-allocation on the happy path by deferring error construction.
// Deprecated: Use local implementations within package to avoid interface allocations.
//
// Performance Note for Future Architect:
// ConvertIndex remains exported and unchanged to protect public API stability.
// In a future major release (v2.0), this function should be removed entirely
// from the public surface or redesigned to return an (int, bool) signature.
// Passing "nameOfIndex string" forces interface allocation overhead during
// string formatting, which introduces an unnecessary performance tax on what
// should be a pure-mathematical execution path.
func ConvertIndex[T any](slice []T, index int, nameOfIndex string) (int, error) {
	newIndex := convertIndex(len(slice), index)

	if newIndex == -1 {
		return -1, fmt.Errorf("%s: %d out of range", nameOfIndex, index)
	}

	return newIndex, nil
}

// convertIndex processes index boundaries at 0 allocs/op on success paths.
// It returns -1 for any out-of-bounds index, allowing callers to handle errors without triggering allocations.
// It handles negative indices by offsetting them from the slice length, mimicking JavaScript's behavior.
func convertIndex(sliceLen, index int) int {
	if index < 0 {
		index += sliceLen
	}
	if index < 0 || index >= sliceLen {
		return -1
	}
	return index
}

// convertEndIndex processes end indices at 0 allocs/op on success paths.
// It returns -1 for any out-of-bounds index, allowing callers to handle errors without triggering allocations.
// It handles negative indices by offsetting them from the slice length, mimicking JavaScript's behavior.
func convertEndIndex(sliceLen, index int) int {
	if index < 0 {
		index += sliceLen
	}

	if index < 0 || index > sliceLen {
		return -1
	}
	return index
}

// isComparable reports whether element can be compared with == without
// panicking. It checks the value, not just its type: an array of slices, or a
// struct whose interface field holds a slice, has a comparable-looking kind but
// makes == panic.
//
// If element passes, == against any other value of type T is safe too: at each
// interface inside it, either the other value holds a different dynamic type
// (== is false) or the same comparable one.
func isComparable[T any](element T) bool {
	var val any = element
	if val == nil {
		return true
	}

	t := reflect.TypeOf(val)
	if !t.Comparable() {
		return false // slices, maps, funcs, and arrays or structs containing them
	}

	switch t.Kind() {
	case reflect.Struct, reflect.Array:
		// The type allows ==, but an interface field or element inside it may
		// hold a slice at run time, so check the value. This costs one
		// allocation, which plain types like int and string skip.
		return reflect.ValueOf(val).Comparable()
	default:
		return true
	}
}

// flatProcessor is a recursive helper function that processes elements for the Flat() method.
// It handles both fast paths (known slice types) and a slow path (reflection) for unexpected slice types.
// The function appends valid T elements to the result slice while respecting the specified depth limit.
func flatProcessor[T any](result *[]T, currentSlice []any, currentDepth int) error {
	sliceLen := len(currentSlice)

	for i := 0; i < sliceLen; i++ {
		v := currentSlice[i]
		if v == nil {
			continue
		}

		// 1. DEPTH CHECK: If depth limit is hit, we CANNOT unpack further. Everything must be T.
		if currentDepth <= 0 {
			if t, ok := v.(T); ok {
				*result = append(*result, t)
				continue
			}
			return fmt.Errorf("element of type %T cannot be converted to T", v)
		}

		// 2. FAST PATHS: We have depth > 0. We MUST try to unpack slices FIRST.
		switch val := v.(type) {
		case []any:
			if err := flatProcessor(result, val, currentDepth-1); err != nil {
				return err
			}
			continue
		case []T:
			// If T is concrete, these elements cannot be unpacked further. Fast-dump them.
			subLen := len(val)
			for j := 0; j < subLen; j++ {
				*result = append(*result, val[j])
			}
			continue
		}

		// 3. SLOW PATH: Fallback to reflection for arbitrary primitive slices ([]int, etc.)
		rv := reflect.ValueOf(v)
		if rv.Kind() == reflect.Slice {
			rvLen := rv.Len()
			for j := 0; j < rvLen; j++ {
				elem := rv.Index(j).Interface()

				// Fast base-case for slow-path elements to avoid recursive allocation
				if currentDepth-1 == 0 {
					if t, ok := elem.(T); ok {
						*result = append(*result, t)
					} else {
						return fmt.Errorf("element of type %T cannot be converted to T", elem)
					}
				} else {
					// Route back into the processor to handle arbitrary deep nesting
					if err := flatProcessor(result, []any{elem}, currentDepth-1); err != nil {
						return err
					}
				}
			}
			continue
		}

		// 4. BASE CASE: Not a slice at all. Treat as scalar T.
		if t, ok := v.(T); ok {
			*result = append(*result, t)
			continue
		}

		return fmt.Errorf("element of type %T cannot be converted to T", v)
	}

	return nil
}

//========================================================
// JOIN UTILITIES
//========================================================

// joinNumeric processes any numeric slice type with 1 alloc/op via compile-time specialization
func joinNumeric[V numeric](s []V, sep string) string {
	var b strings.Builder
	b.Grow(len(s)*4 + len(sep)*(len(s)-1))

	var buf [64]byte
	b.Write(appendNumeric(buf[:0], s[0]))
	for i := 1; i < len(s); i++ {
		b.WriteString(sep)
		b.Write(appendNumeric(buf[:0], s[i]))
	}
	return b.String()
}

// appendNumeric acts as your centralized "numberToString" configuration center
func appendNumeric[V numeric](buf []byte, val V) []byte {
	switch v := any(val).(type) {
	case int:
		return strconv.AppendInt(buf, int64(v), 10)
	case int8:
		return strconv.AppendInt(buf, int64(v), 10)
	case int16:
		return strconv.AppendInt(buf, int64(v), 10)
	case int32:
		return strconv.AppendInt(buf, int64(v), 10)
	case int64:
		return strconv.AppendInt(buf, v, 10)
	case uint:
		return strconv.AppendUint(buf, uint64(v), 10)
	case uint8:
		return strconv.AppendUint(buf, uint64(v), 10)
	case uint16:
		return strconv.AppendUint(buf, uint64(v), 10)
	case uint32:
		return strconv.AppendUint(buf, uint64(v), 10)
	case uint64:
		return strconv.AppendUint(buf, v, 10)
	case float32:
		return strconv.AppendFloat(buf, float64(v), 'g', -1, 32)
	case float64:
		return strconv.AppendFloat(buf, v, 'g', -1, 64)
	default:
		return buf
	}
}

// joinBool leverages strconv.FormatBool's zero-allocation constants
func joinBool(s []bool, sep string) string {
	var b strings.Builder
	b.Grow(len(s)*5 + len(sep)*(len(s)-1))
	b.WriteString(strconv.FormatBool(s[0]))
	for i := 1; i < len(s); i++ {
		b.WriteString(sep)
		b.WriteString(strconv.FormatBool(s[i]))
	}
	return b.String()
}

// joinFallback cleanly handles unoptimized custom structural objects
func joinFallback[T any](slice []T, sep string) string {
	var b strings.Builder
	b.Grow(len(slice)*8 + len(sep)*(len(slice)-1))

	writeFallbackElement(&b, slice[0])
	for i := 1; i < len(slice); i++ {
		b.WriteString(sep)
		writeFallbackElement(&b, slice[i])
	}
	return b.String()
}

func writeFallbackElement[T any](b *strings.Builder, val T) {
	var buf [64]byte
	switch v := any(val).(type) {
	case string:
		b.WriteString(v)
	case bool:
		b.WriteString(strconv.FormatBool(v))

	case int:
		b.Write(appendNumeric(buf[:0], v))
	case int8:
		b.Write(appendNumeric(buf[:0], v))
	case int16:
		b.Write(appendNumeric(buf[:0], v))
	case int32:
		b.Write(appendNumeric(buf[:0], v))
	case int64:
		b.Write(appendNumeric(buf[:0], v))
	case uint:
		b.Write(appendNumeric(buf[:0], v))
	case uint8:
		b.Write(appendNumeric(buf[:0], v))
	case uint16:
		b.Write(appendNumeric(buf[:0], v))
	case uint32:
		b.Write(appendNumeric(buf[:0], v))
	case uint64:
		b.Write(appendNumeric(buf[:0], v))
	case float32:
		b.Write(appendNumeric(buf[:0], v))
	case float64:
		b.Write(appendNumeric(buf[:0], v))

	default:
		b.WriteString(fmt.Sprint(v))
	}
}
