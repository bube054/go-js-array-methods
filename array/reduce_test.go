package array

import (
	"testing"
)

func TestReduce(t *testing.T) {
	tests := []struct {
		name         string
		slice        []int
		initialValue any
		expected     any
		expectError  bool
	}{
		{
			name:         "reduce with initial value",
			slice:        []int{175, 50, 25},
			initialValue: 0,
			expected:     -250, // 0 - 175 - 50 - 25
			expectError:  false,
		},
		{
			name:         "reduce without initial value",
			slice:        []int{1, 2, 3, 4},
			initialValue: nil,
			expected:     -8, // 1 - 2 - 3 - 4
			expectError:  false,
		},
		{
			name:         "reduce empty array without initial errors",
			slice:        []int{},
			initialValue: nil,
			expected:     nil,
			expectError:  true,
		},
		{
			name:         "reduce empty array with initial",
			slice:        []int{},
			initialValue: 42,
			expected:     42,
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myFunc := func(acc any, el int, _ int, _ []int) any {
				accNum := acc.(int)
				return accNum - el
			}

			result, err := Reduce(tt.slice, myFunc, tt.initialValue)

			if tt.expectError && err == nil {
				t.Errorf("Reduce() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("Reduce() got unexpected error: %v", err)
			} else if !tt.expectError && result.(int) != tt.expected.(int) {
				t.Errorf("Reduce() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestReduceStrict(t *testing.T) {
	tests := []struct {
		name        string
		slice       []int
		initialPtr  *int
		expected    int
		expectError bool
	}{
		{
			name:        "reduceStrict without initial",
			slice:       []int{175, 50, 25},
			initialPtr:  nil,
			expected:    100, // 175 - 50 - 25
			expectError: false,
		},
		{
			name:        "reduceStrict with initial",
			slice:       []int{1, 2, 3},
			initialPtr:  func() *int { v := 100; return &v }(),
			expected:    94, // 100 - 1 - 2 - 3
			expectError: false,
		},
		{
			name:        "reduceStrict empty array without initial errors",
			slice:       []int{},
			initialPtr:  nil,
			expected:    0,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			myFunc := func(acc int, el int, _ int, _ []int) int {
				return acc - el
			}

			result, err := ReduceStrict(tt.slice, myFunc, tt.initialPtr)

			if tt.expectError && err == nil {
				t.Errorf("ReduceStrict() expected error but got none")
			} else if !tt.expectError && err != nil {
				t.Errorf("ReduceStrict() got unexpected error: %v", err)
			} else if !tt.expectError && result != tt.expected {
				t.Errorf("ReduceStrict() = %d, expected %d", result, tt.expected)
			}
		})
	}
}

// TestArrayReduce tests the Array.Reduce receiver method
func TestArrayReduce(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4}

	result, err := arr.Reduce(func(acc any, el int, _ int, _ []int) any {
		accNum := acc.(int)
		return accNum + el
	}, 0)

	if err != nil {
		t.Errorf("Array.Reduce() got unexpected error: %v", err)
	}

	expected := 10
	if result.(int) != expected {
		t.Errorf("Array.Reduce() = %v, expected %d", result, expected)
	}
}

// TestArrayReduceStrict tests the Array.ReduceStrict receiver method
func TestArrayReduceStrict(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4}
	initialVal := 0

	result, err := arr.ReduceStrict(func(acc int, el int, _ int, _ []int) int {
		return acc + el
	}, &initialVal)

	if err != nil {
		t.Errorf("Array.ReduceStrict() got unexpected error: %v", err)
	}

	expected := 10
	if result != expected {
		t.Errorf("Array.ReduceStrict() = %d, expected %d", result, expected)
	}
}
