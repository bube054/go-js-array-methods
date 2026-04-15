package array

import (
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected string
	}{
		{
			name:     "toString converts to string",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			expected: "Banana,Orange,Apple,Mango",
		},
		{
			name:     "toString single element",
			slice:    []string{"Hello"},
			expected: "Hello",
		},
		{
			name:     "toString empty array",
			slice:    []string{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToString(tt.slice)
			if result != tt.expected {
				t.Errorf("ToString() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
