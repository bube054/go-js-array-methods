package array

import (
	"fmt"
	"testing"
)

func TestValueOf(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}
	newFruits := ValueOf(fruits)
	fmt.Println("newFruits", newFruits)
}
