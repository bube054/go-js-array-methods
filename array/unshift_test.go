package array

import (
	"fmt"
	"testing"
)

func TestUnShift(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits := UnShift(fruits, "Kiwi", "Lemon")

	fmt.Println("newFruits:", newFruits)
}