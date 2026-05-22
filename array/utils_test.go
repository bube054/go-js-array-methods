package array

import (
	"testing"
)

func TestConvertIndex(t *testing.T) {
	tests := []struct {
		name        string
		index       int
		sliceLength int
		expected    int
		expectError bool
	}{
		{
			name:        "positive index in range",
			index:       2,
			sliceLength: 5,
			expected:    2,
			expectError: false,
		},
		{
			name:        "negative index from end",
			index:       -2,
			sliceLength: 5,
			expected:    3,
			expectError: false,
		},
		{
			name:        "index 0",
			index:       0,
			sliceLength: 5,
			expected:    0,
			expectError: false,
		},
		{
			name:        "index out of bounds positive",
			index:       10,
			sliceLength: 5,
			expected:    -1,
			expectError: true,
		},
		{
			name:        "index out of bounds negative",
			index:       -10,
			sliceLength: 5,
			expected:    -1,
			expectError: true,
		},
		{
			name:        "negative index at end",
			index:       -1,
			sliceLength: 5,
			expected:    4,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary slice for testing
			tempSlice := make([]string, tt.sliceLength)
			result, err := ConvertIndex(tempSlice, tt.index, "index")

			if tt.expectError && err == nil {
				t.Errorf("ConvertIndex() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("ConvertIndex() got unexpected error: %v", err)
			} else if !tt.expectError && result != tt.expected {
				t.Errorf("ConvertIndex() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

func TestConvertIndexBoundaryValues(t *testing.T) {
	sliceLen := 5
	tests := []struct {
		name     string
		index    int
		expected int
		ok       bool
	}{
		{"valid positive index", 2, 2, true},
		{"valid negative index", -2, 3, true},
		{"negative index far out of range", -9999, -1, false},
		{"index equal to length", 5, -1, false},
		{"positive index beyond length", 10, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertIndex(sliceLen, tt.index)
			if (got != -1) != tt.ok {
				t.Fatalf("expected ok=%v but got index=%d", tt.ok, got)
			}
			if got != tt.expected {
				t.Fatalf("expected %d but got %d", tt.expected, got)
			}
		})
	}
}

func TestConvertEndIndexBoundaryValues(t *testing.T) {
	sliceLen := 5
	tests := []struct {
		name     string
		index    int
		expected int
		ok       bool
	}{
		{"valid positive end index", 3, 3, true},
		{"valid negative end index", -1, 4, true},
		{"negative end index far out of range", -9999, -1, false},
		{"end index equal to length", 5, 5, true},
		{"end index beyond length", 10, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertEndIndex(sliceLen, tt.index)
			if (got != -1) != tt.ok {
				t.Fatalf("expected ok=%v but got index=%d", tt.ok, got)
			}
			if got != tt.expected {
				t.Fatalf("expected %d but got %d", tt.expected, got)
			}
		})
	}
}
