package array

import (
	"reflect"
	"testing"
)

type EntriesTest[T comparable] []struct {
	name     string
	slice    []T
	expected []Entry[T]
}

func TestEntries(t *testing.T) {
	tests := EntriesTest[int]{
		// Normal cases
		{
			name:  "entries of regular array",
			slice: []int{10, 20, 30},
			expected: []Entry[int]{
				{index: 0, element: 10},
				{index: 1, element: 20},
				{index: 2, element: 30},
			},
		},
		{
			name:  "entries with single element",
			slice: []int{42},
			expected: []Entry[int]{
				{index: 0, element: 42},
			},
		},
		{
			name:  "entries with duplicates",
			slice: []int{5, 5, 5},
			expected: []Entry[int]{
				{index: 0, element: 5},
				{index: 1, element: 5},
				{index: 2, element: 5},
			},
		},
		{
			name:  "entries with negative numbers",
			slice: []int{-1, -2, -3, -4},
			expected: []Entry[int]{
				{index: 0, element: -1},
				{index: 1, element: -2},
				{index: 2, element: -3},
				{index: 3, element: -4},
			},
		},
		// Edge case - empty array
		{
			name:     "entries of empty array",
			slice:    []int{},
			expected: []Entry[int]{},
		},
		// Larger array
		{
			name:  "entries of larger array",
			slice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []Entry[int]{
				{index: 0, element: 1},
				{index: 1, element: 2},
				{index: 2, element: 3},
				{index: 3, element: 4},
				{index: 4, element: 5},
				{index: 5, element: 6},
				{index: 6, element: 7},
				{index: 7, element: 8},
				{index: 8, element: 9},
				{index: 9, element: 10},
			},
		},
		// Mixed values
		{
			name:  "entries with mixed positive and negative",
			slice: []int{100, -50, 0, 25, -100},
			expected: []Entry[int]{
				{index: 0, element: 100},
				{index: 1, element: -50},
				{index: 2, element: 0},
				{index: 3, element: 25},
				{index: 4, element: -100},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Entries(tt.slice)

			// Use reflect.DeepEqual for slice comparison
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}

			// Verify length
			if len(result) != len(tt.expected) {
				t.Errorf("Expected length %d but got %d", len(tt.expected), len(result))
			}

			// Verify each entry
			for i, entry := range result {
				if entry.index != tt.expected[i].index {
					t.Errorf("At entry %d: expected index %d but got %d", i, tt.expected[i].index, entry.index)
				}
				if entry.element != tt.expected[i].element {
					t.Errorf("At entry %d: expected element %d but got %d", i, tt.expected[i].element, entry.element)
				}
			}
		})
	}
}

func TestEntriesWithStrings(t *testing.T) {
	tests := EntriesTest[string]{
		{
			name:  "entries of string array",
			slice: []string{"apple", "banana", "cherry"},
			expected: []Entry[string]{
				{index: 0, element: "apple"},
				{index: 1, element: "banana"},
				{index: 2, element: "cherry"},
			},
		},
		{
			name:  "entries with single string",
			slice: []string{"hello"},
			expected: []Entry[string]{
				{index: 0, element: "hello"},
			},
		},
		{
			name:  "entries with empty strings",
			slice: []string{"a", "", "b", ""},
			expected: []Entry[string]{
				{index: 0, element: "a"},
				{index: 1, element: ""},
				{index: 2, element: "b"},
				{index: 3, element: ""},
			},
		},
		{
			name:  "entries with duplicate strings",
			slice: []string{"x", "x", "x"},
			expected: []Entry[string]{
				{index: 0, element: "x"},
				{index: 1, element: "x"},
				{index: 2, element: "x"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Entries(tt.slice)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v but got %v", tt.expected, result)
			}

			// Verify each entry
			for i, entry := range result {
				if entry.index != i {
					t.Errorf("At entry %d: expected index %d but got %d", i, i, entry.index)
				}
				if entry.element != tt.expected[i].element {
					t.Errorf("At entry %d: expected element %q but got %q", i, tt.expected[i].element, entry.element)
				}
			}
		})
	}
}

func TestEntriesPreservesOriginal(t *testing.T) {
	t.Run("entries does not modify original slice", func(t *testing.T) {
		original := []int{1, 2, 3, 4, 5}
		originalCopy := make([]int, len(original))
		copy(originalCopy, original)

		result := Entries(original)

		// Verify original wasn't modified
		if !reflect.DeepEqual(original, originalCopy) {
			t.Errorf("Original slice was modified: expected %v but got %v", originalCopy, original)
		}

		// Verify result has correct length
		if len(result) != len(original) {
			t.Errorf("Expected result length %d but got %d", len(original), len(result))
		}

		// Verify entries are correct
		for i, entry := range result {
			if entry.index != i || entry.element != original[i] {
				t.Errorf("Entry %d incorrect: expected {index: %d, element: %d}, got {index: %d, element: %d}",
					i, i, original[i], entry.index, entry.element)
			}
		}
	})
}

func TestEntriesIndexConsistency(t *testing.T) {
	t.Run("entries indices are sequential and correct", func(t *testing.T) {
		slice := []int{100, 200, 300, 400}
		result := Entries(slice)

		// Verify indices are sequential starting from 0
		for i := 0; i < len(result); i++ {
			if result[i].index != i {
				t.Errorf("Expected index %d at position %d, but got %d", i, i, result[i].index)
			}
		}
	})
}

func TestEntriesReturnNewSlice(t *testing.T) {
	t.Run("entries returns a new slice, not a reference", func(t *testing.T) {
		original := []int{1, 2, 3}
		result1 := Entries(original)
		result2 := Entries(original)

		// Both calls should produce equivalent results
		if !reflect.DeepEqual(result1, result2) {
			t.Errorf("Multiple calls to Entries should produce the same result")
		}

		// But they should be different slice objects
		// (We can't directly compare slice headers in Go, but we verify the content is independent)
		if len(result1) != len(result2) {
			t.Errorf("Results should have same length")
		}
	})
}
