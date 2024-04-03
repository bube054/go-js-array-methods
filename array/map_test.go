package array

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T){
	numbers := []int{65, 44, 12, 4}

	newNums := Map(numbers, func(e int, i int, s []int) any {
		return e * 10
	})

	fmt.Println("newNums:", newNums)

	newNums2 := MapStrict(numbers, func(e int, i int, s []int) int {
		return e / 10
	})

	fmt.Println("newNums2:", newNums2)
}