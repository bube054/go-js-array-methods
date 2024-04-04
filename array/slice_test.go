package array

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Lemon", "Apple", "Mango"}
	newFruits, err := Slice(fruits, -3, -1)

	fmt.Println("RES:", newFruits)
	fmt.Println("ERR:", err)
}
