package array

import (
	"reflect"
	"testing"
)

type CopyWithinTest[T comparable] []struct {
	name        string
	slice       []T
	target      int
	start       int
	end         int
	expected    []T
	expectError bool
}

func TestCopyWithin(t *testing.T) {
	tests := CopyWithinTest[int]{
		{
			name:        "copy single element from end to start",
			slice:       []int{1, 2, 3, 4, 5},
			target:      0,
			start:       4,
			end:         5,
			expected:    []int{5, 2, 3, 4, 5},
			expectError: false,
		},
		{
			name:        "copy two elements to start",
			slice:       []int{1, 2, 3, 4, 5},
			target:      0,
			start:       3,
			end:         5,
			expected:    []int{4, 5, 3, 4, 5},
			expectError: false,
		},
		{
			name:        "start greater than end error",
			slice:       []int{1, 2, 3, 4, 5},
			target:      0,
			start:       3,
			end:         1,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CopyWithin(tt.slice, tt.target, tt.start, tt.end)

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error but got: %v", err)
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Expected %v but got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestCopyWithinStrings(t *testing.T) {
	tests := CopyWithinTest[string]{
		{
			name:        "copy strings",
			slice:       []string{"a", "b", "c", "d", "e"},
			target:      0,
			start:       3,
			end:         5,
			expected:    []string{"d", "e", "c", "d", "e"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := CopyWithin(tt.slice, tt.target, tt.start, tt.end)

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error but got: %v", err)
				}

				if !reflect.DeepEqual(result, tt.expected) {
					t.Errorf("Expected %v but got %v", tt.expected, result)
				}
			}
		})
	}
}

// TestArrayCopyWithin tests the Array.CopyWithin receiver method
func TestArrayCopyWithin(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	result, err := arr.CopyWithin(0, 3, 5)

	if err != nil {
		t.Errorf("Array.CopyWithin() got unexpected error: %v", err)
	}

	expected := Array[int]{4, 5, 3, 4, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.CopyWithin() = %v, expected %v", result, expected)
	}
}
