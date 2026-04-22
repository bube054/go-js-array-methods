package array

import (
	"reflect"
	"testing"
)

func TestShift(t *testing.T) {
	tests := []struct {
		name             string
		input            []int
		expectedSlice    []int
		expectedShifted  *int
		expectedShiftNil bool
	}{
		{
			name:             "shift from non-empty array",
			input:            []int{1, 2, 3, 4},
			expectedSlice:    []int{2, 3, 4},
			expectedShifted:  func() *int { v := 1; return &v }(),
			expectedShiftNil: false,
		},
		{
			name:             "shift from single element array",
			input:            []int{42},
			expectedSlice:    []int{},
			expectedShifted:  func() *int { v := 42; return &v }(),
			expectedShiftNil: false,
		},
		{
			name:             "shift from empty array",
			input:            []int{},
			expectedSlice:    []int{},
			expectedShifted:  nil,
			expectedShiftNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, shifted := Shift(tt.input)

			if !reflect.DeepEqual(result, tt.expectedSlice) {
				t.Errorf("Shift() slice = %v, expected %v", result, tt.expectedSlice)
			}

			if tt.expectedShiftNil && shifted != nil {
				t.Errorf("Expected shifted to be nil but got %v", shifted)
			} else if !tt.expectedShiftNil && shifted == nil {
				t.Errorf("Expected shifted to not be nil")
			} else if !tt.expectedShiftNil && *shifted != *tt.expectedShifted {
				t.Errorf("Shift() shifted = %d, expected %d", *shifted, *tt.expectedShifted)
			}
		})
	}
}

func TestShiftStrings(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	newFruits, fruit := Shift(fruits)

	expectedFruits := []string{"Orange", "Apple", "Mango"}
	if !reflect.DeepEqual(newFruits, expectedFruits) {
		t.Errorf("Shift() = %v, expected %v", newFruits, expectedFruits)
	}

	if *fruit != "Banana" {
		t.Errorf("Shifted = %s, expected %s", *fruit, "Banana")
	}
}

// TestArrayShift tests the Array.Shift receiver method
func TestArrayShift(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, shifted := arr.Shift()

	expectedSlice := Array[int]{2, 3, 4, 5}
	expectedShifted := 1

	if !reflect.DeepEqual(result, expectedSlice) {
		t.Errorf("Array.Shift() slice = %v, expected %v", result, expectedSlice)
	}

	if shifted == nil || *shifted != expectedShifted {
		t.Errorf("Array.Shift() shifted = %v, expected %d", shifted, expectedShifted)
	}
}
