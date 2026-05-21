package array

import (
	"testing"
)

func TestSome(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int, int, []int) bool
		expected  bool
	}{
		{
			name:  "some elements > 10",
			slice: []int{3, 10, 18, 20},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: true,
		},
		{
			name:  "some returns false for no match",
			slice: []int{1, 2, 3},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: false,
		},
		{
			name:  "some in empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: false,
		},
		{
			name:  "some with single match",
			slice: []int{100},
			predicate: func(el, _ int, _ []int) bool {
				return el > 50
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Some(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Some() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestArraySome tests the Array.Some receiver method
func TestArraySome(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	predicate := func(el, _ int, _ []int) bool { return el > 3 }

	result := arr.Some(predicate)
	if !result {
		t.Errorf("Array.Some() = %v, expected true", result)
	}
}
