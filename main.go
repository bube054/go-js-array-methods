package main

import (
	"fmt"

	"github.com/bube054/go-js-array-methods/array"
)

func main() {
	a := []int{1, 2, 3, 4}
	r, e := array.Fill[int](a, 0, 2, 3)
	// r, e := array.Fill[int](a, 0, 2, 4)

	fmt.Println(r, e)
}