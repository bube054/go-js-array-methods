package array

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	fruits := []string{}

	res := Reverse(fruits)
	fmt.Println("RES:", res)
}
