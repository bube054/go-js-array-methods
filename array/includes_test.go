package array

import (
	"fmt"
	"testing"
)

func TestIncludes(t *testing.T){
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	start := 1
	present := Includes(fruits, "Orange", &start)

	fmt.Println("Present:", present)
}