package array_test

import (
	"slices"
	"testing"

	"github.com/bube054/go-js-array-methods/v2/array"
)

// -----------------------------------------------------------------------------
// Functional Style Tests
// -----------------------------------------------------------------------------

func TestKeys(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []int
	}{
		{"populated slice", []string{"a", "b", "c"}, []int{0, 1, 2}},
		{"single element", []string{"a"}, []int{0}},
		{"empty slice", []string{}, []int{}},
		{"nil slice", nil, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Go 1.23 slices.Collect consumes the iter.Seq into a standard slice
			result := slices.Collect(array.Keys(tt.input))

			if len(result) == 0 && len(tt.expected) == 0 {
				return // slices.Equal panics/fails on nil vs empty sometimes, safe exit
			}
			if !slices.Equal(result, tt.expected) {
				t.Errorf("Keys() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// -----------------------------------------------------------------------------
// Object-Oriented Style (Array[T]) Tests
// -----------------------------------------------------------------------------

func TestArray_Keys(t *testing.T) {
	arr := array.Array[string]{"x", "y", "z"}
	expected := []int{0, 1, 2}

	result := slices.Collect(arr.Keys())
	if !slices.Equal(result, expected) {
		t.Errorf("Array.Keys() = %v, want %v", result, expected)
	}
}

// -----------------------------------------------------------------------------
// Coverage & Memory Leak Defenses (Short-Circuiting)
// -----------------------------------------------------------------------------

func TestKeys_ShortCircuit(t *testing.T) {
	// This ensures the `if !yield(i) { return }` branch is fully covered.
	input := []string{"a", "b", "c", "d"}
	iterations := 0

	for i := range array.Keys(input) {
		iterations++
		if i == 1 {
			break // Early exit triggers yield() to return false
		}
	}

	if iterations != 2 {
		t.Errorf("Expected iterator to short-circuit after 2 iterations, ran %d times", iterations)
	}
}
