package array

import (
	"fmt"
	"testing"
) 

func TestLastIndexOf(t *testing.T) {
	fruits := []string{"Orange", "Apple", "Mango", "Apple", "Banana", "Apple"}

	start := 3
	present := LastIndexOf(fruits, "Apple", &start)

	fmt.Println("Present:", present)
}
