package array

import (
	"testing"
)

func TestIncludes(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		element  string
		start    *int
		expected bool
	}{
		{
			name:     "includes element",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Orange",
			start:    nil,
			expected: true,
		},
		{
			name:     "includes element from start position",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Orange",
			start:    func() *int { v := 1; return &v }(),
			expected: true,
		},
		{
			name:     "doesn't include element after start",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Banana",
			start:    func() *int { v := 1; return &v }(),
			expected: false,
		},
		{
			name:     "includes element with negative start",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Mango",
			start:    func() *int { v := -1; return &v }(),
			expected: true,
		},
		{
			name:     "doesn't include non-existent element",
			slice:    []string{"Banana", "Orange", "Apple"},
			element:  "Grape",
			start:    nil,
			expected: false,
		},
		{
			name:     "includes in empty array",
			slice:    []string{},
			element:  "test",
			start:    nil,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Includes(tt.slice, tt.element, tt.start)
			if result != tt.expected {
				t.Errorf("Includes() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestIncludesInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	if !Includes(numbers, 3, nil) {
		t.Errorf("Includes(numbers, 3) should be true")
	}
	if Includes(numbers, 10, nil) {
		t.Errorf("Includes(numbers, 10) should be false")
	}
}
