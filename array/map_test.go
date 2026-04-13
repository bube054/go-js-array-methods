package array

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []any
	}{
		{
			name:     "map multiply by 10",
			slice:    []int{65, 44, 12, 4},
			expected: []any{650, 440, 120, 40},
		},
		{
			name:     "map on empty array",
			slice:    []int{},
			expected: []any{},
		},
		{
			name:     "map single element",
			slice:    []int{5},
			expected: []any{50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.slice, func(e int, _ int, _ []int) any {
				return e * 10
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Map() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestMapStrict(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "mapStrict divide by 10",
			slice:    []int{65, 44, 12, 4},
			expected: []int{6, 4, 1, 0},
		},
		{
			name:     "mapStrict on empty array",
			slice:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MapStrict(tt.slice, func(e int, _ int, _ []int) int {
				return e / 10
			})
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("MapStrict() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
