package array

import (
	"testing"
)

func TestLastIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		start    *int
		expected int
	}{
		{
			name:     "lastIndexOf finds last occurrence",
			slice:    []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"},
			element:  "Apple",
			start:    nil,
			expected: 5,
		},
		{
			name:     "lastIndexOf with start position limits search",
			slice:    []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"},
			element:  "Apple",
			start:    func() *int { v := 4; return &v }(),
			expected: 3,
		},
		{
			name:     "lastIndexOf returns -1 for non-existent",
			slice:    []string{"Orange", "Apple", "Mango"},
			element:  "Grape",
			start:    nil,
			expected: -1,
		},
		{
			name:     "lastIndexOf in empty array",
			slice:    []string{},
			element:  "test",
			start:    nil,
			expected: -1,
		},
		{
			name:     "lastIndexOf with start beyond end returns -1",
			slice:    []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"},
			element:  "Apple",
			start:    func() *int { v := 100; return &v }(),
			expected: -1,
		},
		{
			name:     "lastIndexOf with extreme negative start returns -1",
			slice:    []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"},
			element:  "Apple",
			start:    func() *int { v := -9999; return &v }(),
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LastIndexOf(tt.slice, tt.element, tt.start)
			if result != tt.expected {
				t.Errorf("LastIndexOf() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestLastIndexOfInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 3, 5}
	if LastIndexOf(numbers, 3, nil) != 4 {
		t.Errorf("LastIndexOf should find last occurrence at index 4")
	}
	if LastIndexOf(numbers, 10, nil) != -1 {
		t.Errorf("LastIndexOf should return -1 for non-existent")
	}
}

// TestArrayLastIndexOf tests the Array.LastIndexOf receiver method
func TestArrayLastIndexOf(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry", "banana"}

	result := arr.LastIndexOf("banana", nil)
	expected := 3

	if result != expected {
		t.Errorf("Array.LastIndexOf() = %d, expected %d", result, expected)
	}
}

func TestArrayLastIndexOfEmptySlice(t *testing.T) {
	arr := Array[string]{}
	if arr.LastIndexOf("apple", nil) != -1 {
		t.Fatal("Array.LastIndexOf() on empty array should return -1")
	}
}

func TestArrayLastIndexOfInvalidStart(t *testing.T) {
	tests := []struct {
		name     string
		start    int
		expected int
	}{
		{name: "negative start out of bounds", start: -9999, expected: -1},
		{name: "positive start out of bounds", start: 9999, expected: -1},
	}

	arr := Array[string]{"apple", "banana", "cherry", "banana"}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arr.LastIndexOf("banana", &tt.start); got != tt.expected {
				t.Fatalf("Array.LastIndexOf() = %d, expected %d", got, tt.expected)
			}
		})
	}
}

func TestArrayLastIndexOfNonComparable(t *testing.T) {
	arr := Array[[]int]{{1, 2}, {3, 4}, {3, 4}}
	if arr.LastIndexOf([]int{3, 4}, nil) != 2 {
		t.Fatalf("Array.LastIndexOf() = %d, expected %d", arr.LastIndexOf([]int{3, 4}, nil), 2)
	}
	if arr.LastIndexOf([]int{7, 8}, nil) != -1 {
		t.Fatalf("Array.LastIndexOf() = %d, expected -1", arr.LastIndexOf([]int{7, 8}, nil))
	}
}
