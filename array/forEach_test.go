package array

import (
	"fmt"
	"testing"
)

func TestForEach(t *testing.T) {
	numbers := []int{65, 44, 12, 4}

	sum := 0
	ForEach(numbers, func(el int, ind int, list []int) {
		sum += el
	})

	fmt.Println("Sum:", sum)
}
