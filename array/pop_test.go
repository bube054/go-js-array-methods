package array

import (
	"fmt"
	"testing"
)

func TestPop(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	newFruits, fruit := Pop(fruits)

	fmt.Println("newFruits:", newFruits)
	fmt.Println("fruit:", *fruit)
}