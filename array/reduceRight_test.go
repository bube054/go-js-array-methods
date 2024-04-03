package array

import (
	"fmt"
	"testing"
)

func TestReduceRight(t *testing.T) {
	numbers := []int{2, 45, 30, 100}

	myFunc := func(acc any, el int, ind int, slice []int) any {
		accNum := acc.(int)

		return accNum - el
	}

	// init := 0
	res, err := ReduceRight(numbers, myFunc, nil)

	fmt.Println("RES:", res)
	fmt.Println("ERR:", err)
}

func TestReduceRightStrict(t *testing.T) {
	numbers := []int{175, 50, 25}

	myFunc := func(acc int, el int, ind int, slice []int) int {
		return acc - el
	}

	// init := 0
	res, err := ReduceRightStrict(numbers, myFunc, nil)

	fmt.Println("RES:", res)
	fmt.Println("ERR:", err)
}