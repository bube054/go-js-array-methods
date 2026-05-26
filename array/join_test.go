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

func TestArrayJoinInt64(t *testing.T) {
	slice := Array[int64]{1, 2, 3}
	result := slice.Join(",")
	if result != "1,2,3" {
		t.Fatalf("Array.Join() = %s, expected %s", result, "1,2,3")
	}
}

func TestJoinPrimitiveSlices(t *testing.T) {
	tests := []struct {
		name      string
		slice     any
		separator string
		expected  string
	}{
		{name: "bool slice", slice: []bool{true, false, true}, separator: "|", expected: "true|false|true"},
		{name: "int slice", slice: []int{1, -2, 0}, separator: ",", expected: "1,-2,0"},
		{name: "int8 slice", slice: []int8{1, -2}, separator: "-", expected: "1--2"},
		{name: "int16 slice", slice: []int16{300, -300}, separator: ":", expected: "300:-300"},
		{name: "int32 slice", slice: []int32{2147483647, -2147483648}, separator: ";", expected: "2147483647;-2147483648"},
		{name: "int64 slice", slice: []int64{1, 2}, separator: ",", expected: "1,2"},
		{name: "uint slice", slice: []uint{1, 2}, separator: "|", expected: "1|2"},
		{name: "uint8 slice", slice: []uint8{0, 255}, separator: ",", expected: "0,255"},
		{name: "uint16 slice", slice: []uint16{65535, 0}, separator: "|", expected: "65535|0"},
		{name: "uint32 slice", slice: []uint32{4294967295, 1}, separator: "-", expected: "4294967295-1"},
		{name: "uint64 slice", slice: []uint64{1, 2}, separator: ",", expected: "1,2"},
		{name: "float32 slice", slice: []float32{42, 1.5e-7}, separator: "|", expected: "42|1.5e-07"},
		{name: "float64 slice", slice: []float64{123.456789, 1.23e20}, separator: ",", expected: "123.456789,1.23e+20"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sep := tt.separator
			switch s := tt.slice.(type) {
			case []bool:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []int:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []int8:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []int16:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []int32:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []int64:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []uint:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []uint8:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []uint16:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []uint32:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []uint64:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []float32:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			case []float64:
				if got := Join(s, &sep); got != tt.expected {
					t.Fatalf("Join() = %q, want %q", got, tt.expected)
				}
			default:
				t.Fatalf("unsupported slice type %T", tt.slice)
			}
		})
	}
}
