package array

import (
	"testing"
)

func TestJoin(t *testing.T) {
	tests := []struct {
		name      string
		slice     []string
		separator *string
		expected  string
	}{
		{
			name:      "join with custom separator",
			slice:     []string{"Banana", "Orange", "Apple", "Mango"},
			separator: func() *string { s := " and "; return &s }(),
			expected:  "Banana and Orange and Apple and Mango",
		},
		{
			name:      "join with comma separator",
			slice:     []string{"Banana", "Orange", "Apple"},
			separator: func() *string { s := ","; return &s }(),
			expected:  "Banana,Orange,Apple",
		},
		{
			name:      "join with nil separator (default comma)",
			slice:     []string{"Banana", "Orange", "Apple"},
			separator: nil,
			expected:  "Banana,Orange,Apple",
		},
		{
			name:      "join empty array",
			slice:     []string{},
			separator: func() *string { s := ","; return &s }(),
			expected:  "",
		},
		{
			name:      "join single element",
			slice:     []string{"Apple"},
			separator: func() *string { s := ","; return &s }(),
			expected:  "Apple",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Join(tt.slice, tt.separator)
			if result != tt.expected {
				t.Errorf("Join() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

// TestArrayJoin tests the Array.Join standalone function with string slices
func TestArrayJoin(t *testing.T) {
	tests := []struct {
		name      string
		slice     Array[string]
		separator string
		expected  string
	}{
		{
			name:      "Join with custom separator",
			slice:     Array[string]{"apple", "banana", "cherry"},
			separator: " - ",
			expected:  "apple - banana - cherry",
		},
		{
			name:      "Join with comma separator",
			slice:     Array[string]{"apple", "banana", "cherry"},
			separator: ",",
			expected:  "apple,banana,cherry",
		},
		{
			name:      "Join with empty separator",
			slice:     Array[string]{"a", "b", "c"},
			separator: "",
			expected:  "abc",
		},
		{
			name:      "Join with single element",
			slice:     Array[string]{"apple"},
			separator: ",",
			expected:  "apple",
		},
		{
			name:      "Join with empty array",
			slice:     Array[string]{},
			separator: ",",
			expected:  "",
		},
		{
			name:      "Join with space separator",
			slice:     Array[string]{"hello", "world", "go"},
			separator: " ",
			expected:  "hello world go",
		},
		{
			name:      "Join with pipe separator",
			slice:     Array[string]{"red", "green", "blue"},
			separator: "|",
			expected:  "red|green|blue",
		},
		{
			name:      "Join with special characters in separator",
			slice:     Array[string]{"one", "two", "three"},
			separator: ":::",
			expected:  "one:::two:::three",
		},
		{
			name:      "Join with numbers as strings",
			slice:     Array[string]{"1", "2", "3", "4", "5"},
			separator: "-",
			expected:  "1-2-3-4-5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sep := tt.separator
			result := tt.slice.Join(sep)
			if result != tt.expected {
				t.Errorf("Join() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestArrayJoinNonStringTypes(t *testing.T) {
	tests := []struct {
		name      string
		slice     Array[any]
		separator string
		expected  string
	}{
		{
			name:      "Join ints with default separator",
			slice:     Array[any]{1, 2, 3},
			separator: ",",
			expected:  "1,2,3",
		},
		{
			name:      "Join bools with custom separator",
			slice:     Array[any]{true, false, true},
			separator: "|",
			expected:  "true|false|true",
		},
		{
			name:      "Join mixed values",
			slice:     Array[any]{1, "two", 3.0},
			separator: "-",
			expected:  "1-two-3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.slice.Join(tt.separator)
			if result != tt.expected {
				t.Errorf("Join() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
