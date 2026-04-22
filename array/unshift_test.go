package array

import (
	"reflect"
	"testing"
)

func TestUnShift(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		elements []int
		expected []int
	}{
		{
			name:     "unshift single element to non-empty array",
			input:    []int{2, 3},
			elements: []int{1},
			expected: []int{1, 2, 3},
		},
		{
			name:     "unshift multiple elements",
			input:    []int{3, 4},
			elements: []int{1, 2},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "unshift to empty array",
			input:    []int{},
			elements: []int{1, 2},
			expected: []int{1, 2},
		},
		{
			name:     "unshift no elements",
			input:    []int{1, 2},
			elements: []int{},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := UnShift(tt.input, tt.elements...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("UnShift() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestUnShiftStrings(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	newFruits := UnShift(fruits, "Kiwi", "Lemon")

	expected := []string{"Kiwi", "Lemon", "Banana", "Orange", "Apple", "Mango"}
	if !reflect.DeepEqual(newFruits, expected) {
		t.Errorf("UnShift() = %v, expected %v", newFruits, expected)
	}
}

// TestArrayUnShift tests the Array.UnShift receiver method
func TestArrayUnShift(t *testing.T) {
	arr := Array[int]{3, 4, 5}
	result := arr.UnShift(1, 2)

	expected := Array[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.UnShift() = %v, expected %v", result, expected)
	}
}
