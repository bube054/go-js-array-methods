package array

import (
	"fmt"
	"testing"
)

func TestJoin(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	delim := " and "
	newFruits := Join(fruits, &delim)
	fmt.Println("newFruits:", newFruits)
}