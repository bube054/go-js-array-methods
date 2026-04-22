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
