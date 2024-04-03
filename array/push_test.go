package array

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits := Push(fruits, "Kiwi", "Lemon")

	fmt.Println("newFruits:", newFruits)
}