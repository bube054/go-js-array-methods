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
			expected:    []string{"Banana", "Orange", "Kiwi"},
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
		{
			name:        "splice with nil howMany removes to end",
			slice:       []string{"Banana", "Orange", "Apple", "Mango"},
			index:       1,
			howMany:     nil,
			elements:    []string{"Pear"},
			expected:    []string{"Banana", "Pear"},
			expectError: false,
		},
		{
			name:        "splice with negative howMany treats as zero",
			slice:       []string{"Banana", "Orange", "Apple"},
			index:       1,
			howMany:     func() *int { v := -1; return &v }(),
			elements:    []string{"Pear"},
			expected:    []string{"Banana", "Pear", "Orange", "Apple"},
			expectError: false,
		},
		{
			name:        "splice clamps overlarge howMany to end",
			slice:       []string{"Banana", "Orange", "Apple", "Mango"},
			index:       2,
			howMany:     func() *int { v := 10; return &v }(),
			elements:    []string{"Pear"},
			expected:    []string{"Banana", "Orange", "Pear"},
			expectError: false,
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

// TestArraySplice tests the Array.Splice receiver method
func TestArraySplice(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}
	howMany := 2
	result, err := arr.Splice(1, &howMany, 10, 11)

	if err != nil {
		t.Errorf("Array.Splice() got unexpected error: %v", err)
	}

	expectedResult := Array[int]{1, 10, 11, 4, 5}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Array.Splice() result = %v, expected %v", result, expectedResult)
	}
}

func TestSpliceReturnsNonNilEmptySlice(t *testing.T) {
	slice := []int{1}
	howMany := 1
	result, err := Splice(slice, 0, &howMany)

	if err != nil {
		t.Fatalf("expected nil error but got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result for Splice that produces an empty slice")
	}
	if len(result) != 0 {
		t.Fatalf("expected empty result, got length %d", len(result))
	}
}
