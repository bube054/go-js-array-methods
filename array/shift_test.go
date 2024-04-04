package array

import (
	"fmt"
	"testing"
)

func TestShift(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits, fruit := Shift(fruits)

	fmt.Println("newFruits:", newFruits)
	fmt.Println("fruit:", *fruit)
}