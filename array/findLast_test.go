package array

import (
	"testing"
)

func TestFindLast(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int, int, []int) bool
		expected  *int
	}{
		{
			name:  "find last element >= 18",
			slice: []int{3, 7, 10, 15, 18, 20, 25},
			predicate: func(el, _ int, _ []int) bool {
				return el >= 18
			},
			expected: func() *int { v := 25; return &v }(),
		},
		{
			name:  "find last of multiple matches",
			slice: []int{5, 12, 8, 130, 44},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: func() *int { v := 44; return &v }(),
		},
		{
			name:  "find last returns nil for no match",
			slice: []int{1, 2, 3},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: nil,
		},
		{
			name:  "find last in empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FindLast(tt.slice, tt.predicate)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("FindLast() = %v, expected %v", result, tt.expected)
			} else if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("FindLast() = %v, expected %v", *result, *tt.expected)
			}
		})
	}
}

// TestArrayFindLast tests the Array.FindLast receiver method
func TestArrayFindLast(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40, 20}
	predicate := func(el, _ int, _ []int) bool { return el == 20 }

	result := arr.FindLast(predicate)
	expected := 20

	if result == nil || *result != expected {
		t.Errorf("Array.FindLast() = %v, expected %d", result, expected)
	}
}
