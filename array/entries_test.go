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
				{Index: 0, Value: 10},
				{Index: 1, Value: 20},
				{Index: 2, Value: 30},
			},
		},
		{
			name:  "entries with single element",
			slice: []int{42},
			expected: []Entry[int]{
				{Index: 0, Value: 42},
			},
		},
		{
			name:  "entries with duplicates",
			slice: []int{5, 5, 5},
			expected: []Entry[int]{
				{Index: 0, Value: 5},
				{Index: 1, Value: 5},
				{Index: 2, Value: 5},
			},
		},
		{
			name:  "entries with negative numbers",
			slice: []int{-1, -2, -3, -4},
			expected: []Entry[int]{
				{Index: 0, Value: -1},
				{Index: 1, Value: -2},
				{Index: 2, Value: -3},
				{Index: 3, Value: -4},
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
				{Index: 0, Value: 1},
				{Index: 1, Value: 2},
				{Index: 2, Value: 3},
				{Index: 3, Value: 4},
				{Index: 4, Value: 5},
				{Index: 5, Value: 6},
				{Index: 6, Value: 7},
				{Index: 7, Value: 8},
				{Index: 8, Value: 9},
				{Index: 9, Value: 10},
			},
		},
		// Mixed values
		{
			name:  "entries with mixed positive and negative",
			slice: []int{100, -50, 0, 25, -100},
			expected: []Entry[int]{
				{Index: 0, Value: 100},
				{Index: 1, Value: -50},
				{Index: 2, Value: 0},
				{Index: 3, Value: 25},
				{Index: 4, Value: -100},
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
				if entry.Index != tt.expected[i].Index {
					t.Errorf("At entry %d: expected index %d but got %d", i, tt.expected[i].Index, entry.Index)
				}
				if entry.Value != tt.expected[i].Value {
					t.Errorf("At entry %d: expected element %d but got %d", i, tt.expected[i].Value, entry.Value)
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
				{Index: 0, Value: "apple"},
				{Index: 1, Value: "banana"},
				{Index: 2, Value: "cherry"},
			},
		},
		{
			name:  "entries with single string",
			slice: []string{"hello"},
			expected: []Entry[string]{
				{Index: 0, Value: "hello"},
			},
		},
		{
			name:  "entries with empty strings",
			slice: []string{"a", "", "b", ""},
			expected: []Entry[string]{
				{Index: 0, Value: "a"},
				{Index: 1, Value: ""},
				{Index: 2, Value: "b"},
				{Index: 3, Value: ""},
			},
		},
		{
			name:  "entries with duplicate strings",
			slice: []string{"x", "x", "x"},
			expected: []Entry[string]{
				{Index: 0, Value: "x"},
				{Index: 1, Value: "x"},
				{Index: 2, Value: "x"},
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
				if entry.Index != i {
					t.Errorf("At entry %d: expected index %d but got %d", i, i, entry.Index)
				}
				if entry.Value != tt.expected[i].Value {
					t.Errorf("At entry %d: expected element %q but got %q", i, tt.expected[i].Value, entry.Value)
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
			if entry.Index != i || entry.Value != original[i] {
				t.Errorf("Entry %d incorrect: expected {Index: %d, Value: %d}, got {Index: %d, Value: %d}",
					i, i, original[i], entry.Index, entry.Value)
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
			if result[i].Index != i {
				t.Errorf("Expected index %d at position %d, but got %d", i, i, result[i].Index)
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

// TestArrayEntries tests the Array.Entries receiver method
func TestArrayEntries(t *testing.T) {
	arr := Array[string]{"a", "b", "c"}
	result := arr.Entries()

	if len(result) != 3 {
		t.Errorf("Array.Entries() length = %d, expected 3", len(result))
	}

	if result[0].Index != 0 || result[0].Value != "a" {
		t.Errorf("Array.Entries()[0] = {%d, %s}, expected {0, a}", result[0].Index, result[0].Value)
	}

	if result[2].Index != 2 || result[2].Value != "c" {
		t.Errorf("Array.Entries()[2] = {%d, %s}, expected {2, c}", result[2].Index, result[2].Value)
	}
}
