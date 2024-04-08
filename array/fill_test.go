package array

import (
	"fmt"
	"testing"
)

func TestFill(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits, err := Fill(fruits, "Kiwi", 1, 2)

	fmt.Println("newFruits:", newFruits)
	fmt.Println("err:", err)
}