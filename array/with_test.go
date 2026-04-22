package array

import (
	"reflect"
	"testing"
)

func TestWith(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		index       int
		value       string
		expected    []string
		expectError bool
	}{
		{
			name:        "with update element at index",
			slice:       []string{"Januar", "Februar", "Mar", "April"},
			index:       2,
			value:       "March",
			expected:    []string{"Januar", "Februar", "March", "April"},
			expectError: false,
		},
		{
			name:        "with negative index",
			slice:       []string{"Januar", "Februar", "Mar", "April"},
			index:       -1,
			value:       "April Updated",
			expected:    []string{"Januar", "Februar", "Mar", "April Updated"},
			expectError: false,
		},
		{
			name:        "with out of bounds index errors",
			slice:       []string{"a", "b", "c"},
			index:       10,
			value:       "x",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := With(tt.slice, tt.index, tt.value)

			if tt.expectError && err == nil {
				t.Errorf("With() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("With() got unexpected error: %v", err)
			} else if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("With() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestArrayWith tests the Array.With receiver method
func TestArrayWith(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.With(2, 99)

	if err != nil {
		t.Errorf("Array.With() got unexpected error: %v", err)
	}

	expected := Array[int]{1, 2, 99, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.With() = %v, expected %v", result, expected)
	}
}
