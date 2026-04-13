package array

import (
	"reflect"
	"testing"
)

func TestPop(t *testing.T) {
	tests := []struct {
		name           string
		input          []int
		expectedSlice  []int
		expectedPopped *int
		expectedPopNil bool
	}{
		{
			name:           "pop from non-empty array",
			input:          []int{1, 2, 3, 4},
			expectedSlice:  []int{1, 2, 3},
			expectedPopped: func() *int { v := 4; return &v }(),
			expectedPopNil: false,
		},
		{
			name:           "pop from single element array",
			input:          []int{42},
			expectedSlice:  []int{},
			expectedPopped: func() *int { v := 42; return &v }(),
			expectedPopNil: false,
		},
		{
			name:           "pop from empty array",
			input:          []int{},
			expectedSlice:  []int{},
			expectedPopped: nil,
			expectedPopNil: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, popped := Pop(tt.input)

			if !reflect.DeepEqual(result, tt.expectedSlice) {
				t.Errorf("Pop() slice = %v, expected %v", result, tt.expectedSlice)
			}

			if tt.expectedPopNil && popped != nil {
				t.Errorf("Expected popped to be nil but got %v", popped)
			} else if !tt.expectedPopNil && popped == nil {
				t.Errorf("Expected popped to not be nil")
			} else if !tt.expectedPopNil && *popped != *tt.expectedPopped {
				t.Errorf("Pop() popped = %d, expected %d", *popped, *tt.expectedPopped)
			}
		})
	}
}

func TestPopStrings(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	newFruits, fruit := Pop(fruits)

	expectedFruits := []string{"Banana", "Orange", "Apple"}
	if !reflect.DeepEqual(newFruits, expectedFruits) {
		t.Errorf("Pop() = %v, expected %v", newFruits, expectedFruits)
	}

	if *fruit != "Mango" {
		t.Errorf("Popped = %s, expected %s", *fruit, "Mango")
	}
}
