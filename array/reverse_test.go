package array

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fruits := []string{"Cecilie", "Lone", "Emil", "Tobias", "Linus"}

	res := Reverse(fruits)
	fmt.Println("RES:", res)
}
