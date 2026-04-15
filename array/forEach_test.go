package array

import (
	"testing"
)

func TestForEach(t *testing.T) {
	// Test with sum accumulation
	numbers := []int{65, 44, 12, 4}
	sum := 0

	ForEach(numbers, func(el int, _ int, _ []int) {
		sum += el
	})

	expectedSum := 125
	if sum != expectedSum {
		t.Errorf("ForEach sum = %d, expected %d", sum, expectedSum)
	}
}

func TestForEachCallback(t *testing.T) {
	tests := []struct {
		name           string
		slice          []string
		callbackCounts int
	}{
		{
			name:           "forEach on non-empty array calls callback for each",
			slice:          []string{"a", "b", "c"},
			callbackCounts: 3,
		},
		{
			name:           "forEach on empty array doesn't call callback",
			slice:          []string{},
			callbackCounts: 0,
		},
		{
			name:           "forEach on single element calls callback once",
			slice:          []string{"x"},
			callbackCounts: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			callCount := 0
			ForEach(tt.slice, func(_ string, _ int, _ []string) {
				callCount++
			})

			if callCount != tt.callbackCounts {
				t.Errorf("callback called %d times, expected %d", callCount, tt.callbackCounts)
			}
		})
	}
}
