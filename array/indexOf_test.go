package array

import (
	"fmt"
	"testing"
) 

func TestIndexOf(t *testing.T) {
	fruits := []string{"Banana", "Orange", "Apple", "Mango"}

	// start := 1
	present := IndexOf(fruits, "Orange", nil)

	fmt.Println("Present:", present)
}
