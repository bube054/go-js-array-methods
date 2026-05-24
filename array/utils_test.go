package array

import (
	"reflect"
	"testing"
)

type flatProcessorTestStruct struct {
	Name  string
	Value int
}

func makeDeepNestedSlice(depth int, value any) []any {
	current := value
	for i := 0; i < depth; i++ {
		current = []any{current}
	}
	return []any{current}
}

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
func TestOptionalParam(t *testing.T) {
	if OptionalParam([]int{10, 20}, 5) != 10 {
		t.Fatalf("OptionalParam([]int{10,20},5) = %d, expected 10", OptionalParam([]int{10, 20}, 5))
	}
	if OptionalParam([]int{}, 5) != 5 {
		t.Fatalf("OptionalParam([]int{},5) = %d, expected 5", OptionalParam([]int{}, 5))
	}
}

func TestIsComparable(t *testing.T) {
	if !isComparable(10) {
		t.Fatal("expected int to be comparable")
	}
	if isComparable([]int{1, 2}) {
		t.Fatal("expected slice to be non-comparable")
	}
	var nilAny any
	if !isComparable(nilAny) {
		t.Fatal("expected nil to be comparable")
	}
	if isComparable(map[string]int{"a": 1}) {
		t.Fatal("expected map to be non-comparable")
	}
	if isComparable(func() {}) {
		t.Fatal("expected func to be non-comparable")
	}
	if isComparable(struct{ Items []int }{Items: []int{1}}) {
		t.Fatal("expected struct containing slice to be non-comparable")
	}
}

func TestFlatProcessor_BasicFlattening(t *testing.T) {
	result := make([]any, 0, 8)
	input := []any{1, []any{2, 3}, 4}

	if err := flatProcessor(&result, input, 1); err != nil {
		t.Fatalf("flatProcessor() returned unexpected error: %v", err)
	}

	expected := []any{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("flatProcessor() = %v, want %v", result, expected)
	}
	if cap(result) != 8 {
		t.Fatalf("flatProcessor() cap = %d, want %d", cap(result), 8)
	}
}

func TestFlatProcessor_ConcreteTypeFastPath(t *testing.T) {
	result := make([]int, 0, 8)
	input := []any{[]int{1, 2}, []int{3, 4}}

	if err := flatProcessor(&result, input, 1); err != nil {
		t.Fatalf("flatProcessor() returned unexpected error: %v", err)
	}

	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("flatProcessor() = %v, want %v", result, expected)
	}
	if cap(result) != 8 {
		t.Fatalf("flatProcessor() cap = %d, want %d", cap(result), 8)
	}
}

func TestFlatProcessor_SlowPathRecursiveDepth(t *testing.T) {
	result := make([]any, 0, 4)
	input := []any{[]int{1, 2}}

	if err := flatProcessor(&result, input, 2); err != nil {
		t.Fatalf("flatProcessor() returned unexpected error: %v", err)
	}

	expected := []any{1, 2}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("flatProcessor() = %v, want %v", result, expected)
	}
}

func TestFlatProcessor_SlowPathRecursiveError(t *testing.T) {
	result := make([]string, 0, 2)
	input := []any{[]int{1}}

	err := flatProcessor(&result, input, 2)
	if err == nil {
		t.Fatal("flatProcessor() expected error, got nil")
	}
	if err.Error() != "element of type int cannot be converted to T" {
		t.Fatalf("flatProcessor() error = %q, want %q", err.Error(), "element of type int cannot be converted to T")
	}
}

func TestFlatProcessor_DepthBehavior(t *testing.T) {
	tests := []struct {
		name     string
		input    []any
		depth    int
		expected []any
	}{
		{
			name:     "depth 0 leaves nested slices intact",
			input:    []any{1, []any{2, []any{3, 4}}, 5},
			depth:    0,
			expected: []any{1, []any{2, []any{3, 4}}, 5},
		},
		{
			name:     "depth 1 flattens one level",
			input:    []any{1, []any{2, []any{3, 4}}, 5},
			depth:    1,
			expected: []any{1, 2, []any{3, 4}, 5},
		},
		{
			name:     "depth 2 flattens nested slices fully",
			input:    []any{1, []any{2, []any{3, 4}}, 5},
			depth:    2,
			expected: []any{1, 2, 3, 4, 5},
		},
		{
			name:     "negative depth behaves like depth 0",
			input:    []any{1, []any{2, 3}, 4},
			depth:    -1,
			expected: []any{1, []any{2, 3}, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]any, 0, len(tt.input)+2)
			if err := flatProcessor(&result, tt.input, tt.depth); err != nil {
				t.Fatalf("flatProcessor() returned unexpected error: %v", err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Fatalf("flatProcessor() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFlatProcessor_DeepNesting(t *testing.T) {
	input := makeDeepNestedSlice(25, "deep")
	result := make([]any, 0, 1)

	if err := flatProcessor(&result, input, 25); err != nil {
		t.Fatalf("flatProcessor() returned unexpected error: %v", err)
	}

	expected := []any{"deep"}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("flatProcessor() = %v, want %v", result, expected)
	}
}

func TestFlatProcessor_HeterogeneousTypesAndPreallocation(t *testing.T) {
	value := flatProcessorTestStruct{Name: "nested", Value: 7}
	ptr := &flatProcessorTestStruct{Name: "pointer", Value: 9}
	input := []any{"alpha", nil, []any{42, value, ptr, true}, 3.14}
	result := make([]any, 0, 16)

	if err := flatProcessor(&result, input, 1); err != nil {
		t.Fatalf("flatProcessor() returned unexpected error: %v", err)
	}

	expected := []any{"alpha", 42, value, ptr, true, 3.14}
	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("flatProcessor() = %v, want %v", result, expected)
	}
	if cap(result) != 16 {
		t.Fatalf("flatProcessor() cap = %d, want %d", cap(result), 16)
	}
}

func TestFlatProcessor_Errors(t *testing.T) {
	tests := []struct {
		name     string
		depth    int
		input    []any
		expected string
	}{
		{
			name:     "depth 0 rejects non-matching scalar",
			depth:    0,
			input:    []any{"bad"},
			expected: "element of type string cannot be converted to T",
		},
		{
			name:     "base case rejects non-matching scalar",
			depth:    1,
			input:    []any{"bad"},
			expected: "element of type string cannot be converted to T",
		},
		{
			name:     "slow path rejects non-matching element",
			depth:    1,
			input:    []any{[]string{"bad"}},
			expected: "element of type string cannot be converted to T",
		},
		{
			name:     "nested recursion propagates error",
			depth:    1,
			input:    []any{[]any{"bad"}},
			expected: "element of type string cannot be converted to T",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]int, 0, 4)
			err := flatProcessor(&result, tt.input, tt.depth)
			if err == nil {
				t.Fatalf("flatProcessor() expected error, got nil")
			}
			if err.Error() != tt.expected {
				t.Fatalf("flatProcessor() error = %q, want %q", err.Error(), tt.expected)
			}
		})
	}
}
