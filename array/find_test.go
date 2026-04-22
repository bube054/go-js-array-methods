package array

import (
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int, int, []int) bool
		expected  *int
	}{
		{
			name:  "find first element > 18",
			slice: []int{3, 10, 18, 20},
			predicate: func(el, _ int, _ []int) bool {
				return el > 18
			},
			expected: func() *int { v := 20; return &v }(),
		},
		{
			name:  "find no match returns nil",
			slice: []int{1, 2, 3},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: nil,
		},
		{
			name:  "find in empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: nil,
		},
		{
			name:  "find first match from beginning",
			slice: []int{5, 12, 8, 130, 44},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: func() *int { v := 12; return &v }(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Find(tt.slice, tt.predicate)
			if (result == nil && tt.expected != nil) || (result != nil && tt.expected == nil) {
				t.Errorf("Find() = %v, expected %v", result, tt.expected)
			} else if result != nil && tt.expected != nil && *result != *tt.expected {
				t.Errorf("Find() = %v, expected %v", *result, *tt.expected)
			}
		})
	}
}

// TestArrayFind tests the Array.Find receiver method
func TestArrayFind(t *testing.T) {
	arr := Array[int]{10, 20, 30, 40}
	predicate := func(el, _ int, _ []int) bool { return el > 25 }

	result := arr.Find(predicate)
	expected := 30

	if result == nil || *result != expected {
		t.Errorf("Array.Find() = %v, expected %d", result, expected)
	}
}
