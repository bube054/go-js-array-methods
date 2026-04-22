package array

import (
	"reflect"
	"testing"
)

type ConcatTestCase[T comparable] struct {
	name     string
	slice1   []T
	slice2   [][]T
	expected []T
}

func TestConcat(t *testing.T) {
	tests := []ConcatTestCase[int]{
		// Normal cases
		{
			name:     "concat two non-empty slices",
			slice1:   []int{1, 2, 3},
			slice2:   [][]int{{4, 5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "concat three slices",
			slice1:   []int{1, 2},
			slice2:   [][]int{{3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "concat with empty first slice",
			slice1:   []int{},
			slice2:   [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "concat with empty second slice",
			slice1:   []int{1, 2, 3},
			slice2:   [][]int{{}, {4, 5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "concat all empty slices",
			slice1:   []int{},
			slice2:   [][]int{{}, {}},
			expected: []int{},
		},
		// Edge cases
		{
			name:     "single element in each slice",
			slice1:   []int{1},
			slice2:   [][]int{{2}, {3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "only first slice has elements",
			slice1:   []int{1, 2, 3},
			slice2:   [][]int{{}, {}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "concat multiple slices",
			slice1:   []int{1},
			slice2:   [][]int{{2}, {3}, {4}, {5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "concat single element to many",
			slice1:   []int{1, 2, 3, 4, 5},
			slice2:   [][]int{{6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "concat many elements to single",
			slice1:   []int{1},
			slice2:   [][]int{{2, 3, 4, 5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Concat(tt.slice1, tt.slice2...)

			// Verify length
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d but got %d", len(tt.expected), len(result))
			}

			// Use reflect.DeepEqual for slice comparison
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
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

func TestConcatWithStrings(t *testing.T) {
	tests := []ConcatTestCase[string]{
		{
			name:     "concat string slices",
			slice1:   []string{"a", "b"},
			slice2:   [][]string{{"c", "d"}, {"e"}},
			expected: []string{"a", "b", "c", "d", "e"},
		},
		{
			name:     "concat single strings",
			slice1:   []string{"hello"},
			slice2:   [][]string{{"world"}, {"!"}},
			expected: []string{"hello", "world", "!"},
		},
		{
			name:     "concat with empty string elements",
			slice1:   []string{"a", ""},
			slice2:   [][]string{{"", "b"}},
			expected: []string{"a", "", "", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Concat(tt.slice1, tt.slice2...)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}
		})
	}
}

func TestConcatPreservesOriginal(t *testing.T) {
	t.Run("concat does not modify original slices", func(t *testing.T) {
		original1 := []int{1, 2, 3}
		original2 := []int{4, 5, 6}
		original1Copy := make([]int, len(original1))
		original2Copy := make([]int, len(original2))
		copy(original1Copy, original1)
		copy(original2Copy, original2)

		result := Concat(original1, original2)

		// Verify originals weren't modified
		if !reflect.DeepEqual(original1, original1Copy) {
			t.Errorf("original1 was modified: expected %v but got %v", original1Copy, original1)
		}
		if !reflect.DeepEqual(original2, original2Copy) {
			t.Errorf("original2 was modified: expected %v but got %v", original2Copy, original2)
		}

		// Verify result is correct
		expected := []int{1, 2, 3, 4, 5, 6}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})
}

// TestArrayConcat tests the Array.Concat receiver method
func TestArrayConcat(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	slice2 := Array[int]{4, 5, 6}
	slice3 := Array[int]{7, 8}

	result := arr.Concat(slice2, slice3)
	expected := Array[int]{1, 2, 3, 4, 5, 6, 7, 8}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Concat() = %v, expected %v", result, expected)
	}
}
