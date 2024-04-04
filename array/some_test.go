package array

import (
	"fmt"
	"testing"
)

func TestSome(t *testing.T){
	ages := []int{3, 10, 18, 20}

	age := Some(ages, func(age int, ind int, list []int) bool {
		return age >  100
	})

	fmt.Println("RES:", age)
}