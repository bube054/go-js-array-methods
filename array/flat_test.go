package array

import (
	"reflect"
	"testing"
)

func TestFlat(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "Flatten basic 2D array",
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Empty 2D array",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "Single element sub-arrays",
			input:    [][]int{{1}, {2}, {3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Mixed size sub-arrays",
			input:    [][]int{{1}, {2, 3, 4}, {5, 6}},
			expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			name:     "Array with empty sub-array",
			input:    [][]int{{1, 2}, {}, {3, 4}},
			expected: []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Flat(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Flat(%v) length = %d, expected %d", tt.input, len(result), len(tt.expected))
			} else if len(result) > 0 && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Flat(%v) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFlatStrings(t *testing.T) {
	input := [][]string{{"Hello", "World"}, {"Foo", "Bar"}, {"Baz"}}
	expected := []string{"Hello", "World", "Foo", "Bar", "Baz"}

	result := Flat(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Flat(%v) = %v, expected %v", input, result, expected)
	}
}

func TestFlatEmptySubArrays(t *testing.T) {
	input := [][]int{{}, {}, {}}
	expected := []int{}

	result := Flat(input)

	if len(result) != len(expected) {
		t.Errorf("Flat(%v) length = %d, expected %d", input, len(result), len(expected))
	}
}
