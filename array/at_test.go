package array

import (
	"fmt"
	"testing"
)

func TestAt(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	want := "Apple"
	index := -5
	result, err := At(fruits, index)

	if err != nil {
		fmt.Println("Err", err)
		t.Fatalf(`At() = %q, %v, want match for %#q, nil`, result, err, want)
	}

	fmt.Println("Result:", result)
}
