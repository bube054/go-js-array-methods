package array

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "reverse array",
			slice:    []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"},
			expected: []string{"Linus", "Tobias", "Emil", "Lone", "Cecilie"},
		},
		{
			name:     "reverse single element",
			slice:    []string{"Hello"},
			expected: []string{"Hello"},
		},
		{
			name:     "reverse empty array",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "reverse two elements",
			slice:    []string{"A", "B"},
			expected: []string{"B", "A"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.slice)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Reverse() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestReverseInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	result := Reverse(numbers)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse() = %v, expected %v", result, expected)
	}
}
