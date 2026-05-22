package array

import (
	"fmt"
	"reflect"
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
