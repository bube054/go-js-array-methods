package array

import (
	"testing"
)

func TestFindIndex(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int, int, []int) bool
		expected  int
	}{
		{
			name:  "find index of first element >= 18",
			slice: []int{3, 10, 18, 20},
			predicate: func(el, _ int, _ []int) bool {
				return el >= 18
			},
			expected: 2,
		},
		{
			name:  "find index returns -1 for no match",
			slice: []int{1, 2, 3},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: -1,
		},
		{
			name:  "find index in empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: -1,
		},
		{
			name:  "find index at start",
			slice: []int{5, 12, 8, 130, 44},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindIndex(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("FindIndex() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
