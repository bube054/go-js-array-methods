package array

import (
	"reflect"
	"testing"
)

func TestSplice(t *testing.T) {
	tests := []struct {
		name        string
		slice       []string
		index       int
		howMany     *int
		elements    []string
		expected    []string
		expectError bool
	}{
		{
			name:        "splice remove 2 elements",
			slice:       []string{"Banana", "Orange", "Apple", "Mango", "Kiwi"},
			index:       2,
			howMany:     func() *int { v := 2; return &v }(),
			elements:    []string{},
			expected:    []string{"Banana", "Orange", "Mango", "Kiwi"},
			expectError: false,
		},
		{
			name:        "splice insert elements",
			slice:       []string{"Banana", "Orange", "Apple", "Mango"},
			index:       0,
			howMany:     func() *int { v := 2; return &v }(),
			elements:    []string{"Mango", "Coconut"},
			expected:    []string{"Mango", "Coconut", "Apple", "Mango"},
			expectError: false,
		},
		{
			name:        "splice with invalid index errors",
			slice:       []string{"a", "b", "c"},
			index:       10,
			howMany:     nil,
			elements:    []string{},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Splice(tt.slice, tt.index, tt.howMany, tt.elements...)

			if tt.expectError && err == nil {
				t.Errorf("Splice() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Splice() got unexpected error: %v", err)
			} else if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Splice() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
