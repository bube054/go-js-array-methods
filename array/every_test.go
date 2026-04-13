package array

import (
	"testing"
)

type EveryTest[T comparable] []struct{
	name      string
		slice     []T
		predicate func(T, int, []T) bool
		expected  bool
}
func TestEvery(t *testing.T) {
	tests := EveryTest[int]{
		{
			name:  "all elements greater than 0",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: true,
		},
		{
			name:  "not all elements greater than 3",
			slice: []int{1, 2, 3, 4, 5},
			predicate: func(el, _ int, _ []int) bool {
				return el > 3
			},
			expected: false,
		},
		{
			name:  "all elements in empty array",
			slice: []int{},
			predicate: func(el, _ int, _ []int) bool {
				return el > 0
			},
			expected: true,
		},
		{
			name:  "single element passes",
			slice: []int{5},
			predicate: func(el, _ int, _ []int) bool {
				return el == 5
			},
			expected: true,
		},
		{
			name:  "single element fails",
			slice: []int{5},
			predicate: func(el, _ int, _ []int) bool {
				return el == 3
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Every(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Every() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestEvery_String(t *testing.T){
	tests := EveryTest[string]{
		{
			name:  "all elements non-empty",
			slice: []string{"a", "b", "c"},
			predicate: func(el string, _ int, _ []string) bool {
				return el != ""
			},
			expected: true,
		},
		{
			name:  "not all elements non-empty",
			slice: []string{"a", "", "c"},
			predicate: func(el string, _ int, _ []string) bool {
				return el != ""
			},
			expected: false,
		},
		{
			name:  "all elements in empty array",
			slice: []string{},

			predicate: func(el string, _ int, _ []string) bool {
				return el != ""
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Every(tt.slice, tt.predicate)
			if result != tt.expected {
				t.Errorf("Every() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
