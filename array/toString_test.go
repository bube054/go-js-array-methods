package array

import (
	"fmt"
	"testing"
)

func TestToString(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	res := ToString(fruits)
	fmt.Println("RES:", res)
}