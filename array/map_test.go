package array

import (
	"testing"
)

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "map multiply by 10",
			slice:    []int{65, 44, 12, 4},
			expected: []int{650, 440, 120, 40},
		},
		{
			name:     "map on empty array",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "map single element",
			slice:    []int{5},
			expected: []int{50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Map(tt.slice, func(e int, _ int, _ []int) any {
				return e * 10
			})
			// Enforce type safety: assert result elements are int
			resultInts := make([]int, len(result))
			for i, v := range result {
				intVal, ok := v.(int)
				if !ok {
					t.Errorf("Map() result[%d] is not int: %v", i, v)
					return
				}
				resultInts[i] = intVal
			}
			// Direct type-safe comparison
			if len(resultInts) != len(tt.expected) {
				t.Errorf("Map() length = %d, expected %d", len(resultInts), len(tt.expected))
			} else {
				for i, v := range resultInts {
					if v != tt.expected[i] {
						t.Errorf("Map() = %v, expected %v", resultInts, tt.expected)
						break
					}
				}
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
			// Enforce type safety: direct comparison for []int
			if len(result) != len(tt.expected) {
				t.Errorf("MapStrict() length = %d, expected %d", len(result), len(tt.expected))
			} else {
				for i, v := range result {
					if v != tt.expected[i] {
						t.Errorf("MapStrict() = %v, expected %v", result, tt.expected)
						break
					}
				}
			}
		})
	}
}
