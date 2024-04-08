package array

import (
	"fmt"
	"testing"
)

func TestAt(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	index := -2
	result, err := At(fruits, index)

	if err != nil {
		fmt.Println("Err", err)
	}

	fmt.Println("Result:", result)
}
