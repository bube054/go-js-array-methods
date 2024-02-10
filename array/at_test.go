package array

import (
	"fmt"
	"testing"
)

func TestAt(t *testing.T) {
	array1 := []int{5, 12, 8, 130, 44}

	r, err := At[int](array1, 2)

	fmt.Println(r, err)

	if *r != 8 || err != nil {
		t.Error("Expected &\"8\", nil, but got", r, err)
	}

	r, err = At[int](array1, -2)

	if *r != 130 || err != nil {
		t.Error("Expected &\"130\", nil, but got", r, err)
	}

	r, err = At[int](array1, 5)

	if err == nil {
		t.Error("Expected error but got", r, err)
	}

	r, err = At[int](array1, -5)

	if err == nil {
		t.Error("Expected error but got", r, err)
	}
}