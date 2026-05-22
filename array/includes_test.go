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
		{
			name:     "includes with start beyond end returns false",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Mango",
			start:    func() *int { v := 100; return &v }(),
			expected: false,
		},
		{
			name:     "includes with extreme negative start returns false",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			element:  "Banana",
			start:    func() *int { v := -9999; return &v }(),
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

// TestArrayIncludes tests the Array.Includes receiver method
func TestArrayIncludes(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry"}

	result := arr.Includes("banana", nil)
	if !result {
		t.Errorf("Array.Includes() = %v, expected true", result)
	}

	result = arr.Includes("grape", nil)
	if result {
		t.Errorf("Array.Includes() = %v, expected false", result)
	}
}

func TestArrayIncludesNonComparableElement(t *testing.T) {
	arr := Array[[]int]{{1, 2}, {3, 4}, {5, 6}}
	if !arr.Includes([]int{3, 4}, nil) {
		t.Errorf("Array.Includes() should return true for deeply equal slice element")
	}
	if arr.Includes([]int{7, 8}, nil) {
		t.Errorf("Array.Includes() should return false for missing slice element")
	}
}

func TestArrayIncludesStartBounds(t *testing.T) {
	arr := Array[string]{"apple", "banana", "cherry"}
	if arr.Includes("banana", func() *int { v := 1; return &v }()) == false {
		t.Fatal("Array.Includes() should return true when start is within bounds")
	}
	if arr.Includes("banana", func() *int { v := 100; return &v }()) {
		t.Fatal("Array.Includes() should return false when start is beyond end")
	}
}
