package array

import (
	"reflect"
	"testing"
)

func TestFill(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		value       string
		start       int
		end         int
		expected    []string
		expectError bool
	}{
		{
			name:        "fill range with value",
			slice:       []string{"Banana", "Orange", "Apple", "Mango"},
			value:       "Kiwi",
			start:       1,
			end:         2,
			expected:    []string{"Banana", "Kiwi", "Apple", "Mango"},
			expectError: false,
		},
		{
			name:        "fill with negative indices",
			slice:       []string{"a", "b", "c", "d"},
			value:       "x",
			start:       -2,
			end:         -1,
			expected:    []string{"a", "b", "x", "d"},
			expectError: false,
		},
		{
			name:        "fill start > end errors",
			slice:       []string{"a", "b", "c"},
			value:       "x",
			start:       3,
			end:         1,
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Fill(tt.slice, tt.value, tt.start, tt.end)

			if tt.expectError && err == nil {
				t.Errorf("Fill() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Fill() got unexpected error: %v", err)
			} else if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Fill() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
