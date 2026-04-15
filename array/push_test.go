// TESTFILECONTENT:
package array

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		elements  []int
		expected  []int
		expectLen int
	}{
		// Normal cases
		{
			name:      "push single element to non-empty slice",
			slice:     []int{1, 2, 3},
			elements:  []int{4},
			expected:  []int{1, 2, 3, 4},
			expectLen: 4,
		},
		{
			name:      "push multiple elements",
			slice:     []int{1, 2},
			elements:  []int{3, 4, 5},
			expected:  []int{1, 2, 3, 4, 5},
			expectLen: 5,
		},
		{
			name:      "push to empty slice",
			slice:     []int{},
			elements:  []int{1, 2, 3},
			expected:  []int{1, 2, 3},
			expectLen: 3,
		},
		// Edge cases
		{
			name:      "push single element to empty slice",
			slice:     []int{},
			elements:  []int{42},
			expected:  []int{42},
			expectLen: 1,
		},
		{
			name:      "push empty list (no elements)",
			slice:     []int{1, 2, 3},
			elements:  []int{},
			expected:  []int{1, 2, 3},
			expectLen: 3,
		},
		{
			name:      "push negative numbers",
			slice:     []int{1, 2},
			elements:  []int{-1, -2, -3},
			expected:  []int{1, 2, -1, -2, -3},
			expectLen: 5,
		},
		{
			name:      "push zero values",
			slice:     []int{1, 2, 3},
			elements:  []int{0, 0, 0},
			expected:  []int{1, 2, 3, 0, 0, 0},
			expectLen: 6,
		},
		// Larger arrays
		{
			name:      "push to large array",
			slice:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			elements:  []int{11, 12},
			expected:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			expectLen: 12,
		},
		{
			name:      "push many elements",
			slice:     []int{1},
			elements:  []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expectLen: 10,
		},
		{
			name:      "push single element to large array",
			slice:     []int{1, 2, 3, 4, 5},
			elements:  []int{6},
			expected:  []int{1, 2, 3, 4, 5, 6},
			expectLen: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Push(tt.slice, tt.elements...)

			// Use reflect.DeepEqual for slice comparison
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}

			// Verify length
			if len(result) != tt.expectLen {
				t.Errorf("Expected length %d but got %d", tt.expectLen, len(result))
			}

			// Verify each element
			for i, val := range result {
				if val != tt.expected[i] {
					t.Errorf("At index %d: expected %d but got %d", i, tt.expected[i], val)
				}
			}
		})
	}
}

func TestPushWithStrings(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		elements []string
		expected []string
	}{
		{
			name:     "push string elements",
			slice:    []string{"apple", "banana"},
			elements: []string{"cherry", "date"},
			expected: []string{"apple", "banana", "cherry", "date"},
		},
		{
			name:     "push single string",
			slice:    []string{"hello"},
			elements: []string{"world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "push empty strings",
			slice:    []string{"a", "b"},
			elements: []string{"", "c"},
			expected: []string{"a", "b", "", "c"},
		},
		{
			name:     "push to empty string array",
			slice:    []string{},
			elements: []string{"first", "second", "third"},
			expected: []string{"first", "second", "third"},
		},
		{
			name:     "push duplicate strings",
			slice:    []string{"x"},
			elements: []string{"x", "x"},
			expected: []string{"x", "x", "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Push(tt.slice, tt.elements...)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}

			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d but got %d", len(tt.expected), len(result))
			}
		})
	}
}

func TestPushOrderPreservation(t *testing.T) {
	t.Run("push preserves order of elements", func(t *testing.T) {
		slice := []int{1, 2, 3}
		elements := []int{4, 5, 6}
		result := Push(slice, elements...)

		// Verify original elements are in correct order
		for i := 0; i < len(slice); i++ {
			if result[i] != slice[i] {
				t.Errorf("Original element at index %d: expected %d but got %d", i, slice[i], result[i])
			}
		}

		// Verify new elements are in correct order
		for i := 0; i < len(elements); i++ {
			if result[len(slice)+i] != elements[i] {
				t.Errorf("New element at index %d: expected %d but got %d", len(slice)+i, elements[i], result[len(slice)+i])
			}
		}
	})
}

func TestPushReturnValue(t *testing.T) {
	t.Run("push returns new slice", func(t *testing.T) {
		slice := []int{1, 2, 3}
		result := Push(slice, 4, 5)

		// Verify result contains all elements
		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}

		// Verify the length is correct
		if len(result) != 5 {
			t.Errorf("Expected length 5 but got %d", len(result))
		}

		// Verify capacity is at least the length
		if cap(result) < len(result) {
			t.Errorf("Capacity should be at least %d but got %d", len(result), cap(result))
		}
	})
}

func TestPushWithVariadicArguments(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		args     []int
		expected []int
	}{
		{
			name:     "variadic with 1 argument",
			slice:    []int{1, 2, 3},
			args:     []int{4},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "variadic with 5 arguments",
			slice:    []int{1},
			args:     []int{2, 3, 4, 5, 6},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "variadic with 10 arguments",
			slice:    []int{},
			args:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Push(tt.slice, tt.args...)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}
		})
	}
}

func TestPushMultipleCalls(t *testing.T) {
	t.Run("multiple sequential push calls", func(t *testing.T) {
		slice := []int{1}
		slice = Push(slice, 2)
		slice = Push(slice, 3, 4)
		slice = Push(slice, 5)

		expected := []int{1, 2, 3, 4, 5}
		if !reflect.DeepEqual(slice, expected) {
			t.Errorf("Expected %v but got %v", expected, slice)
		}
	})
}
