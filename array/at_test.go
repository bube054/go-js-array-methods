package array

import (
	"strings"
	"testing"
)

func TestAtComprehensive(t *testing.T) {
	tests := []struct {
		name          string
		slice         []int
		index         int
		expectedValue int
		expectError   bool
	}{
		// Normal cases - positive index
		{
			name:          "positive index at start",
			slice:         []int{10, 20, 30, 40, 50},
			index:         0,
			expectedValue: 10,
			expectError:   false,
		},
		{
			name:          "positive index in middle",
			slice:         []int{10, 20, 30, 40, 50},
			index:         2,
			expectedValue: 30,
			expectError:   false,
		},
		{
			name:          "positive index at end",
			slice:         []int{10, 20, 30, 40, 50},
			index:         4,
			expectedValue: 50,
			expectError:   false,
		},
		// Normal cases - negative index
		{
			name:          "negative index -1 (last element)",
			slice:         []int{10, 20, 30, 40, 50},
			index:         -1,
			expectedValue: 50,
			expectError:   false,
		},
		{
			name:          "negative index -2 (second to last)",
			slice:         []int{10, 20, 30, 40, 50},
			index:         -2,
			expectedValue: 40,
			expectError:   false,
		},
		{
			name:          "negative index -5 (first element)",
			slice:         []int{10, 20, 30, 40, 50},
			index:         -5,
			expectedValue: 10,
			expectError:   false,
		},
		// Edge cases
		{
			name:          "single element array positive index",
			slice:         []int{42},
			index:         0,
			expectedValue: 42,
			expectError:   false,
		},
		{
			name:          "single element array negative index",
			slice:         []int{42},
			index:         -1,
			expectedValue: 42,
			expectError:   false,
		},
		// Out of bounds - positive index
		{
			name:          "out of bounds positive index",
			slice:         []int{10, 20, 30},
			index:         5,
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "out of bounds positive index exactly at length",
			slice:         []int{10, 20, 30},
			index:         3,
			expectedValue: 0,
			expectError:   true,
		},
		// Out of bounds - negative index
		{
			name:          "out of bounds negative index",
			slice:         []int{10, 20, 30},
			index:         -10,
			expectedValue: 0,
			expectError:   true,
		},
		{
			name:          "out of bounds negative index exact boundary",
			slice:         []int{10, 20, 30},
			index:         -4,
			expectedValue: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := At(tt.slice, tt.index)

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
				if !strings.Contains(err.Error(), "out of range") {
					t.Errorf("Expected error containing 'out of range', got '%v'", err)
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error but got: %v", err)
				}
				if result != tt.expectedValue {
					t.Errorf("Expected value %d but got %d", tt.expectedValue, result)
				}
			}
		})
	}
}

func TestAtWithStrings(t *testing.T) {
	tests := []struct {
		name          string
		slice         []string
		index         int
		expectedValue string
		expectError   bool
	}{
		{
			name:          "string array positive index",
			slice:         []string{"apple", "banana", "cherry"},
			index:         1,
			expectedValue: "banana",
			expectError:   false,
		},
		{
			name:          "string array negative index",
			slice:         []string{"apple", "banana", "cherry"},
			index:         -1,
			expectedValue: "cherry",
			expectError:   false,
		},
		{
			name:          "string array out of bounds",
			slice:         []string{"apple", "banana", "cherry"},
			index:         5,
			expectedValue: "",
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := At(tt.slice, tt.index)

			if tt.expectError {
				if err == nil {
					t.Fatalf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Fatalf("Expected no error but got: %v", err)
				}
				if result != tt.expectedValue {
					t.Errorf("Expected value %q but got %q", tt.expectedValue, result)
				}
			}
		})
	}
}
