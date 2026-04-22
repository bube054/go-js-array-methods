package array

import (
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "valueOf returns slice itself",
			slice:    []string{"Banana", "Orange", "Apple", "Mango"},
			expected: []string{"Banana", "Orange", "Apple", "Mango"},
		},
		{
			name:     "valueOf single element",
			slice:    []string{"Hello"},
			expected: []string{"Hello"},
		},
		{
			name:     "valueOf empty array",
			slice:    []string{},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValueOf(tt.slice)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ValueOf() = %v, expected %v", result, tt.expected)
			}
		})
	}
}

func TestValueOfInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := ValueOf(numbers)
	if !reflect.DeepEqual(result, numbers) {
		t.Errorf("ValueOf() should return the same slice")
	}
}

// TestArrayValueOf tests the Array.ValueOf receiver method
func TestArrayValueOf(t *testing.T) {
	arr := Array[int]{1, 2, 3}
	result := arr.ValueOf()

	expected := Array[int]{1, 2, 3}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Array.ValueOf() = %v, expected %v", result, expected)
	}
}
