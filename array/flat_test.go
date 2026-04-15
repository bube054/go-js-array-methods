package array

import (
	"reflect"
	"testing"
)

// TestFlat tests the Flat function
func TestFlat(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		depths    []int
		expected  []int
		expectErr bool
	}{
		{
			name: "Flat with default depth (1)",
			input: []any{
				1,
				[]any{2, 3},
				4,
			},
			depths:    []int{},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name: "Flat with depth 2",
			input: []any{
				1,
				[]any{2, []any{3, 4}},
				5,
			},
			depths:    []int{2},
			expected:  []int{1, 2, 3, 4, 5},
			expectErr: false,
		},
		{
			name: "Flat with depth 0 (no flattening)",
			input: []any{
				1,
				2,
				3,
				4,
			},
			depths:    []int{0},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name: "Flat with deeply nested arrays and sufficient depth",
			input: []any{
				1,
				[]any{
					2,
					[]any{
						3,
						[]any{4},
					},
				},
			},
			depths:    []int{3},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name: "Flat with depth greater than nesting",
			input: []any{
				1,
				[]any{2, 3},
				[]any{4},
			},
			depths:    []int{10},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name: "Flat with empty nested arrays",
			input: []any{
				1,
				[]any{},
				2,
				[]any{3},
			},
			depths:    []int{1},
			expected:  []int{1, 2, 3},
			expectErr: false,
		},
		{
			name: "Flat with single element",
			input: []any{
				42,
			},
			depths:    []int{1},
			expected:  []int{42},
			expectErr: false,
		},
		{
			name:      "Flat with empty array",
			input:     []any{},
			depths:    []int{1},
			expected:  nil,
			expectErr: false,
		},
		{
			name: "Flat with negative depth",
			input: []any{
				1,
				2,
				3,
				4,
			},
			depths:    []int{-1},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
		{
			name: "Flat with partial nesting (depth 1 on 2-level nesting)",
			input: []any{
				1,
				[]any{2, []any{3, 4}},
				5,
			},
			depths:    []int{1},
			expected:  nil,
			expectErr: true, // Error because nested []any can't be converted to int
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Flat[int](tt.input, tt.depths...)
			if (err != nil) != tt.expectErr {
				t.Errorf("Flat() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Flat() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestFlatErrors tests error cases for the Flat function
func TestFlatErrors(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		depths    []int
		expectErr bool
	}{
		{
			name: "Flat with incompatible type at depth 0",
			input: []any{
				"string",
				123,
			},
			depths:    []int{0},
			expectErr: true,
		},
		{
			name: "Flat with mixed types at flattening depth",
			input: []any{
				1,
				[]any{"string", 2},
				3,
			},
			depths:    []int{1},
			expectErr: true,
		},
		{
			name: "Flat with nested arrays that can't be flattened to target type",
			input: []any{
				1,
				[]any{2, []any{3, 4}},
			},
			depths:    []int{1},
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Flat[int](tt.input, tt.depths...)
			if (err != nil) != tt.expectErr {
				t.Errorf("Flat() error = %v, expectErr %v", err, tt.expectErr)
			}
		})
	}
}

// TestFlatWithTypedSlices tests the Flat function with typed slices (not []any)
func TestFlatWithTypedSlices(t *testing.T) {
	tests := []struct {
		name      string
		input     []any
		depths    []int
		expected  []int
		expectErr bool
	}{
		{
			name: "Flat with typed []int slice",
			input: []any{
				1,
				2,
				[]int{1, 3, 5},
				4,
			},
			depths:    []int{1},
			expected:  []int{1, 2, 1, 3, 5, 4},
			expectErr: false,
		},
		{
			name: "Flat with multiple typed slices",
			input: []any{
				1,
				[]int{2, 3},
				[]int{4, 5},
				6,
			},
			depths:    []int{1},
			expected:  []int{1, 2, 3, 4, 5, 6},
			expectErr: false,
		},
		{
			name: "Flat with nested typed slices",
			input: []any{
				1,
				[]any{2, []int{3, 4}},
				5,
			},
			depths:    []int{2},
			expected:  []int{1, 2, 3, 4, 5},
			expectErr: false,
		},
		{
			name: "Flat with empty typed slice",
			input: []any{
				1,
				[]int{},
				2,
			},
			depths:    []int{1},
			expected:  []int{1, 2},
			expectErr: false,
		},
		{
			name: "Flat with typed slice at depth 1 only",
			input: []any{
				1,
				[]int{2, 3},
				4,
			},
			depths:    []int{1},
			expected:  []int{1, 2, 3, 4},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Flat[int](tt.input, tt.depths...)
			if (err != nil) != tt.expectErr {
				t.Errorf("Flat() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Flat() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

// TestArrayFlat tests the Array.Flat receiver method
func TestArrayFlat(t *testing.T) {
	tests := []struct {
		name      string
		arr       Array[any]
		depths    []int
		expected  Array[any]
		expectErr bool
	}{
		{
			name: "Array.Flat with default depth",
			arr: Array[any]{
				1,
				[]any{2, 3},
				4,
			},
			depths: []int{},
			expected: Array[any]{
				1,
				2,
				3,
				4,
			},
			expectErr: false,
		},
		{
			name: "Array.Flat with depth 2",
			arr: Array[any]{
				1,
				[]any{2, []any{3, 4}},
				5,
			},
			depths: []int{2},
			expected: Array[any]{
				1,
				2,
				3,
				4,
				5,
			},
			expectErr: false,
		},
		{
			name:      "Array.Flat with single element",
			arr:       Array[any]{1},
			depths:    []int{1},
			expected:  Array[any]{1},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.arr.Flat(tt.depths...)
			if (err != nil) != tt.expectErr {
				t.Errorf("Array.Flat() error = %v, expectErr %v", err, tt.expectErr)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Array.Flat() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestArrayFlatSmoke(t *testing.T) {
	tests := []struct {
		name     string
		arr      Array[any]
		depths   []int
		expected Array[any]
	}{
		{
			name: "Smoke test with typed []int slice",
			arr: Array[any]{
				1,
				[]int{2, 3},
				4,
			},
			depths: []int{1},
			expected: Array[any]{
				1,
				2,
				3,
				4,
			},
		},
		{
			name: "Smoke test with nested typed []int slice",
			arr: Array[any]{
				1,
				[]any{2, []int{3, 4}},
				5,
			},
			depths: []int{2},
			expected: Array[any]{
				1,
				2,
				3,
				4,
				5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.arr.Flat(tt.depths...)
			if err != nil {
				t.Fatalf("Array.Flat() unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Array.Flat() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
