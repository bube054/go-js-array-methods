package array

import (
	"fmt"
	"testing"
)

func TestReduce(t *testing.T) {
	numbers := []int{175, 50, 25}

	myFunc := func(acc any, el int, ind int, slice []int) any {
		accNum := acc.(int)

		return accNum - el
	}

	init := 0
	res, err := Reduce(numbers, myFunc, init)

	fmt.Println("RES:", res)
	fmt.Println("ERR:", err)
}

func TestReduceStrict(t *testing.T) {
	numbers := []int{175, 50, 25}

	myFunc := func(acc int, el int, ind int, slice []int) int {
		return acc - el
	}

	// init := 0
	res, err := ReduceStrict(numbers, myFunc, nil)

	fmt.Println("RES:", res)
	fmt.Println("ERR:", err)
}