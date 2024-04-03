package array

import (
	"testing"
	"fmt"
)

func TestEntries(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango", "Kiwi", "Papaya"}

	newFruits := Entries(fruits)

	fmt.Println("New:", newFruits)
}
