package array

import (
	"testing"
	"fmt"
)

func TestEntries(t *testing.T) {
	fruits := []string{}

	newFruits := Entries(fruits)

	fmt.Println("New:", newFruits)
}
