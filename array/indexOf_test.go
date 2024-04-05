package array

import (
	"fmt"
	"testing"
) 

func TestIndexOf(t *testing.T) {
	fruits := []string{"Orange", "Banana", "Orange", "Apple", "Mango"}

	start := 1
	present := IndexOf(fruits, "Orange", &start)

	fmt.Println("Present:", present)
}
