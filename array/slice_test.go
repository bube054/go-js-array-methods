package array

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		start       int
		end         int
		expected    []string
		expectError bool
	}{
		{
			name:        "slice with negative indices",
			slice:       []string{"Banana", "Orange", "Lemon", "Apple", "Mango"},
			start:       -3,
			end:         -1,
			expected:    []string{"Lemon", "Apple"},
			expectError: false,
		},
		{
			name:        "slice with positive indices",
			slice:       []string{"Banana", "Orange", "Lemon", "Apple", "Mango"},
			start:       1,
			end:         3,
			expected:    []string{"Orange", "Lemon"},
			expectError: false,
		},
		{
			name:        "slice start greater than end errors",
			slice:       []string{"a", "b", "c"},
			start:       2,
			end:         1,
			expected:    nil,
			expectError: true,
		},
		{
			name:        "slice out of bounds start errors",
			slice:       []string{"a", "b", "c"},
			start:       10,
			end:         1,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Slice(tt.slice, tt.start, tt.end)

			if tt.expectError && err == nil {
				t.Errorf("Slice() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Slice() got unexpected error: %v", err)
			} else if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Slice() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestSliceInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result, _ := Slice(numbers, 1, 3)
	expected := []int{2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Slice() = %v, expected %v", result, expected)
	}
}

// TestArraySlice tests the Array.Slice receiver method
func TestArraySlice(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.Slice(1, 4)

	if err != nil {
		t.Errorf("Array.Slice() got unexpected error: %v", err)
	}

	expected := Array[int]{2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.Slice() = %v, expected %v", result, expected)
	}
}
