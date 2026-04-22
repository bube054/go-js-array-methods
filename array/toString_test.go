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

// TestArrayToString tests the Array.ToString receiver method
func TestArrayToString(t *testing.T) {
	tests := []struct {
		name     string
		arr      Array[string]
		expected string
	}{
		{
			name:     "ToString with multiple elements",
			arr:      Array[string]{"Banana", "Orange", "Apple", "Mango"},
			expected: "Banana,Orange,Apple,Mango",
		},
		{
			name:     "ToString with empty array",
			arr:      Array[string]{},
			expected: "",
		},
		{
			name:     "ToString with single element",
			arr:      Array[string]{"Apple"},
			expected: "Apple",
		},
		{
			name:     "ToString with numbers as strings",
			arr:      Array[string]{"1", "2", "3", "4", "5"},
			expected: "1,2,3,4,5",
		},
		{
			name:     "ToString with special characters",
			arr:      Array[string]{"hello", "world", "go!"},
			expected: "hello,world,go!",
		},
		{
			name:     "ToString with spaces",
			arr:      Array[string]{"hello world", "foo bar"},
			expected: "hello world,foo bar",
		},
		{
			name:     "ToString with empty strings",
			arr:      Array[string]{"", "hello", "", "world", ""},
			expected: ",hello,,world,",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.arr.ToString()
			if result != tt.expected {
				t.Errorf("Array.ToString() = %s, expected %s", result, tt.expected)
			}
		})
	}
}
