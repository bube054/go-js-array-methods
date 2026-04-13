package array

import (
	"testing"
)

func TestIndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		start    *int
		expected int
	}{
		{
			name:     "indexOf element at beginning",
			slice:    []string{"Orange", "Banana", "Orange", "Apple", "Mango"},
			element:  "Orange",
			start:    nil,
			expected: 0,
		},
		{
			name:     "indexOf element from start position",
			slice:    []string{"Orange", "Banana", "Orange", "Apple", "Mango"},
			element:  "Orange",
			start:    func() *int { v := 1; return &v }(),
			expected: 2,
		},
		{
			name:     "indexOf returns -1 for non-existent",
			slice:    []string{"Orange", "Banana", "Apple"},
			element:  "Grape",
			start:    nil,
			expected: -1,
		},
		{
			name:     "indexOf in empty array",
			slice:    []string{},
			element:  "test",
			start:    nil,
			expected: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IndexOf(tt.slice, tt.element, tt.start)
			if result != tt.expected {
				t.Errorf("IndexOf() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestIndexOfInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 3}
	if IndexOf(numbers, 3, nil) != 2 {
		t.Errorf("IndexOf should find first occurrence")
	}
	if IndexOf(numbers, 10, nil) != -1 {
		t.Errorf("IndexOf should return -1 for non-existent")
	}
}
