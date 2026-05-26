package array_test

import (
	"slices"
	"testing"

	"github.com/bube054/go-js-array-methods/v2/array"
)

// -----------------------------------------------------------------------------
// Functional Style Tests
// -----------------------------------------------------------------------------

func TestValues(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{"populated slice", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"single element", []string{"a"}, []string{"a"}},
		{"empty slice", []string{}, []string{}},
		{"nil slice", nil, []string{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := slices.Collect(array.Values(tt.input))

			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Values() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Object-Oriented Style (Array[T]) Tests
// -----------------------------------------------------------------------------

func TestArray_Values(t *testing.T) {
	arr := array.Array[string]{"x", "y", "z"}
	expected := []string{"x", "y", "z"}

	result := slices.Collect(arr.Values())
	if !slices.Equal(result, expected) {
		t.Errorf("Array.Values() = %v, want %v", result, expected)
	}
}

// -----------------------------------------------------------------------------
// Coverage & Memory Leak Defenses (Short-Circuiting)
// -----------------------------------------------------------------------------

func TestValues_ShortCircuit(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	iterations := 0

	for v := range array.Values(input) {
		iterations++
		if v == "b" {
			break
		}
	}

	if iterations != 2 {
		t.Errorf("Expected iterator to short-circuit after 2 iterations, ran %d times", iterations)
	}
}
