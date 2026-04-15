package array

import (
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	tests := []struct {
		name      string
		slice     []int
		predicate func(int, int, []int) bool
		expected  []int
	}{
		{
			name:  "filter elements greater than 18",
			slice: []int{32, 33, 16, 40},
			predicate: func(el, _ int, _ []int) bool {
				return el > 18
			},
			expected: []int{32, 33, 40},
		},
		{
			name:  "filter with no matches",
			slice: []int{1, 2, 3},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: []int{},
		},
		{
			name:  "filter with all matches",
			slice: []int{20, 25, 30},
			predicate: func(el, _ int, _ []int) bool {
				return el > 10
			},
			expected: []int{20, 25, 30},
		},
		{
			name:  "filter empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Filter(tt.slice, tt.predicate)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Filter() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
