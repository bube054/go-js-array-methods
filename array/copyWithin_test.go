package array

import (
	"fmt"
	"testing"
)

func TestCopyWithin(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

	newFruits, err := CopyWithin(fruits, -6, -4, 2)

	fmt.Println("New:", newFruits)
	fmt.Println("New:", err)
}
