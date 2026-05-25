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

func isComparable[T any](element T) bool {
	var val any = element
	if val == nil {
		return true
	}

	switch reflect.TypeOf(element).Kind() {
	case reflect.Slice, reflect.Map, reflect.Func:
		return false
	case reflect.Struct:
		return reflect.TypeOf(element).Comparable()
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

// writeAny is a helper function that writes any value to a strings.Builder.
func writeAny[T any](b *strings.Builder, val T) {
	var buf [64]byte // Stack-allocated buffer for primitive parsing
	switch v := any(val).(type) {
	case string:
		b.WriteString(v)
	case int:
		b.Write(strconv.AppendInt(buf[:0], int64(v), 10))
	case int64:
		b.Write(strconv.AppendInt(buf[:0], v, 10))
	case float64:
		b.Write(strconv.AppendFloat(buf[:0], v, 'g', -1, 64))
	case bool:
		b.WriteString(strconv.FormatBool(v))
	default:
		b.WriteString(fmt.Sprint(v))
	}
}
