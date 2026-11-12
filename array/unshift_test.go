package array

import (
	"reflect"
	"testing"
)

func TestUnshift(t *testing.T) {
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
		{
			name:     "unshift nil slice with no elements returns nil",
			input:    nil,
			elements: []int{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Unshift(tt.input, tt.elements...)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Unshift() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestUnshiftStrings(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	newFruits := Unshift(fruits, "Kiwi", "Lemon")

	expected := []string{"Kiwi", "Lemon", "Banana", "Orange", "Apple", "Mango"}
	if !reflect.DeepEqual(newFruits, expected) {
		t.Errorf("Unshift() = %v, expected %v", newFruits, expected)
	}
}

// TestArrayUnshift tests the Array.Unshift receiver method
func TestArrayUnshift(t *testing.T) {
	arr := Array[int]{3, 4, 5}
	result := arr.Unshift(1, 2)

	expected := Array[int]{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Unshift() = %v, expected %v", result, expected)
	}
}

// The deprecated UnShift names must keep working for existing callers.
func TestDeprecatedUnShiftStillWorks(t *testing.T) {
	if got := UnShift([]int{3}, 1, 2); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("UnShift() = %v, expected [1 2 3]", got)
	}
	if got := (Array[int]{3}).UnShift(1, 2); !reflect.DeepEqual(got, Array[int]{1, 2, 3}) {
		t.Errorf("Array.UnShift() = %v, expected [1 2 3]", got)
	}
}
