package array

import (
	"testing"
)

func TestReduceRight(t *testing.T) {
	tests := []struct {
		name         string
		slice        []int
		initialValue int
		expected     int
		expectError  bool
	}{
		{
			name:         "reduceRight with initial value",
			slice:        []int{100},
			initialValue: 300,
			expected:     200, // 300 - 100
			expectError:  false,
		},
		{
			name:         "reduceRight multiple elements",
			slice:        []int{1, 2, 3, 4},
			initialValue: 100,
			expected:     90, // 100 - 4 - 3 - 2 - 1 (right to left)
			expectError:  false,
		},
		{
			name:         "reduceRight empty array without initial errors",
			slice:        []int{},
			initialValue: 0,
			expected:     0,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myFunc := func(acc any, el int, _ int, _ []int) any {
				return acc.(int) - el
			}

			result, err := ReduceRight(tt.slice, myFunc, tt.initialValue)

			if tt.expectError && err == nil {
				t.Errorf("ReduceRight() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("ReduceRight() got unexpected error: %v", err)
			} else if !tt.expectError && result != tt.expected {
				t.Errorf("ReduceRight() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestReduceRightStrict(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		initialPtr  *int
		expected    int
		expectError bool
	}{
		{
			name:        "reduceRightStrict without initial",
			slice:       []int{175, 50, 25},
			initialPtr:  nil,
			expected:    -200, // 25 - 50 - 175 (start with last element 25, then process backward)
			expectError: false,
		},
		{
			name:        "reduceRightStrict with initial",
			slice:       []int{1, 2, 3},
			initialPtr:  func() *int { v := 100; return &v }(),
			expected:    94, // 100 - 3 - 2 - 1 (right to left)
			expectError: false,
		},
		{
			name:        "reduceRightStrict empty array with initial returns initial",
			slice:       []int{},
			initialPtr:  func() *int { v := 50; return &v }(),
			expected:    50,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myFunc := func(acc int, el int, _ int, _ []int) int {
				return acc - el
			}

			result, err := ReduceRightStrict(tt.slice, myFunc, tt.initialPtr)

			if tt.expectError && err == nil {
				t.Errorf("ReduceRightStrict() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("ReduceRightStrict() got unexpected error: %v", err)
			} else if !tt.expectError && result != tt.expected {
				t.Errorf("ReduceRightStrict() = %d, expected %d", result, tt.expected)
			}
		})
	}
}
func TestArrayReduceRightReceiverMethods(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4}

	reduceRightResult, err := arr.ReduceRight(func(acc any, el int, _ int, _ []int) any {
		return acc.(int) - el
	}, 100)
	if err != nil {
		t.Fatalf("Array.ReduceRight() unexpected error: %v", err)
	}
	if reduceRightResult != 90 {
		t.Fatalf("Array.ReduceRight() = %v, expected %v", reduceRightResult, 90)
	}

	initial := 100
	reduceRightStrictResult, err := arr.ReduceRightStrict(func(acc int, el int, _ int, _ []int) int {
		return acc - el
	}, &initial)
	if err != nil {
		t.Fatalf("Array.ReduceRightStrict() unexpected error: %v", err)
	}
	if reduceRightStrictResult != 90 {
		t.Fatalf("Array.ReduceRightStrict() = %d, expected %d", reduceRightStrictResult, 90)
	}
}
