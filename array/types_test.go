package array

import (
	"reflect"
	"testing"
)

// TestArrayMethodChaining verifies that methods returning Array[T] can be chained.
func TestArrayMethodChaining(t *testing.T) {
	arr := Array[int]{1, 2, 3, 4, 5}

	chain1 := arr.Filter(func(el, _ int, _ []int) bool {
		return el%2 == 0
	}).Push(6).Reverse()

	expected1 := Array[int]{6, 4, 2}
	if !reflect.DeepEqual(chain1, expected1) {
		t.Errorf("chained Filter->Push->Reverse = %v, expected %v", chain1, expected1)
	}

	chain2 := arr.Concat(Array[int]{6, 7}) .Reverse()
	expected2 := Array[int]{7, 6, 5, 4, 3, 2, 1}
	if !reflect.DeepEqual(chain2, expected2) {
		t.Errorf("chained Concat->Reverse = %v, expected %v", chain2, expected2)
	}
}
