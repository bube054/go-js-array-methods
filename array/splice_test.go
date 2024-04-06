package array

import (
	"fmt"
	"testing"
)

func TestSplice(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi"}

	index := 2
	res, err := Splice(fruits, 2, &index)
	fmt.Println("RES:", res) // [Banana Orange Mango Kiwi]
	fmt.Println("ERR:", err) // nil

	index = 2
	res, err = Splice(fruits, 0, &index, "Mango", "Coconut")
	fmt.Println("RES:", res) // [Mango Coconut Apple Mango Kiwi]
	fmt.Println("ERR:", err) // nil
}